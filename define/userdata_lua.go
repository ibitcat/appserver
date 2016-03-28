// 用户缓存数据的redis脚本

package define

// ps: G开头表示全局变量

// 更新用户发出的红包
// KEYS = define.UserCachePrefix + userId
// ARGV = define.EUser_Outcome, year, money
var GUser_Outcome = `
	local total = 0 -- 发出的红包总额
	local exists = redis.call('EXISTS',KEYS[1])
	if exists == 1 then -- 用户缓存存在
		local outcomeStr = redis.call('HGET',KEYS[1],ARGV[1])
		local old = {}
		local year = ARGV[2]
		local money = tonumber(ARGV[3])
		if outcomeStr then
			old = cjson.decode(outcomeStr)
			if old and old[ARGV[1]] then --有这一年的发红包记录
				old[year][1] = old[year][1]+ money
				old[year][2] = old[year][2]+ 1
			else
				old[year]= {money,1}
			end
		else
			old[year]= {money,1}
		end

		local jsStr = cjson.encode(old)
		redis.call('HSET',KEYS[1],ARGV[1],jsStr)
		for k,v in pairs(old) do
			total = total + (v[1] or 0)
		end
	end
	return total
`
