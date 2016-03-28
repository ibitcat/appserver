// 抢红包的redis脚本

package define

// ps: G开头表示全局变量
// 另外需要注意：抢红包系统中，GrabStatusKey 与 DeviceKey必定会同时操作

// 抢红包的脚本
// KEYS = this.GrabKey,this.GrabStatusKey,this.DeviceKey
// ARGV = exp, userId, deviceId, now, status
var GRedpktLua_Grab = `
	-- 添加到抢到的用户队列中
	local num = redis.call('ZADD',KEYS[1],ARGV[1],ARGV[2])
	if num==1 then
		--抢红包状态
		local status = {
			status = tonumber(ARGV[5]),
			expire = tonumber(ARGV[1]),
			device = ARGV[3]
		}
		local status_json_str = cjson.encode(status)
		redis.call('HSET',KEYS[2],ARGV[2],status_json_str)
		
		--设备
		local str = ARGV[2] .. ":" .. tostring(ARGV[1])
		redis.call("HSET",KEYS[3], ARGV[3], str)

		return 0
	elseif num == 0 then
		return 10257
	end

	return 10253
`

// 分享红包任务脚本
// KEYS = this.GrabStatusKey, this.GrabKey, this.DeviceKey
// ARGV = userId, exp, now, deviceId, status
var GRedpktLua_Share = `
	-- 抢红包状态
	local statusStr = redis.call('HGET',KEYS[1],ARGV[1])
	if statusStr then
		local statusJson = cjson.decode(statusStr)
		local status = tonumber(statusJson.status)
		local expire = tonumber(statusJson.expire)
		--redis.log(redis.LOG_NOTICE,"------->",statusJson.expire,statusJson.status)
		if status == 1 and tonumber(ARGV[3]) >= expire then 
			return 10256 -- 分享超时
		end 
		
		if status == 2 and tonumber(ARGV[3]) < expire then
			return 10257 -- 已经分享，等待上传截图
		end

		-- 设备是否一致(先判断红包是否抢到)
		local deviceStr = redis.call('HGET',KEYS[3],ARGV[4])
		if deviceStr then
			local userid,exp = string.match(deviceStr,"(%w+):(%w+)")
			if userid ~= ARGV[1] then
				return 10264
			end
		else
			return 10264
		end
	else
		return 10256
	end
	
	-- 分享
	local count = redis.call('ZADD',KEYS[2],ARGV[2],ARGV[1])
	if count==1 then --分享超时 
		redis.call('ZREM',KEYS[2],ARGV[1])
		return 10256
	end

	--抢红包状态
	local newstatus = {
		status = tonumber(ARGV[5]),
		expire = tonumber(ARGV[2]),
		device = ARGV[4]
	}
	local status_json_str = cjson.encode(newstatus)
	redis.call('HSET',KEYS[1],ARGV[1],status_json_str)
	redis.log(redis.LOG_WARNING,"[LuaLog] " .. status_json_str)
	
	--设备
	local str = ARGV[1] .. ":" .. tostring(ARGV[2])
	redis.call("HSET",KEYS[3], ARGV[4], str)
	return 0
`

// 完成红包的脚本
// KEYS = this.GrabKey, this.RecordKey, this.DeviceKey
// ARGV = userId, deviceId, now, total
var GRedpktLua_Finish = `
	-- 红包剩余个数
	local zcard = redis.call('ZCARD',KEYS[1])
	local hlen = redis.call('HLEN',KEYS[2])
	if tonumber(zcard) + tonumber(hlen)>=ARGV[4] then
		return 10254
	end

	-- 设备是否一致
	local deviceStr = redis.call('HGET',KEYS[3],ARGV[2])
	if deviceStr then
		local userid,exp = string.match(deviceStr,"(%w+):(%w+)")
		if userid ~= ARGV[1] then
			return 10264
		end
	else
		return 10264
	end

	local count = redis.call('ZREM',KEYS[1],ARGV[1])
	if count == 0 then
		return 10261
	end
	redis.call('ZADD', KEYS[2], ARGV[3], ARGV[1])

	return 0
`
