local M = {}

M.orc_points = {
	{ pos = vmath.vector3(60, 75, 0), taken = false }, -- Point 1
	{ pos = vmath.vector3(145, 75, 0), taken = false },     -- Point 2
	{ pos = vmath.vector3(240, 75, 0), taken = false },    -- Point 3
	{ pos = vmath.vector3(330, 75, 0), taken = false },    -- Point 4
	{ pos = vmath.vector3(440, 75, 0), taken = false }    -- Point 5
}

-- position = M.orc_point[1].pos
-- taken = M.orc_point[1].taken
function M.set_taken(index, value) -- points.set_taken(1, true)
    if M.orc_points[index] then
        M.orc_points[index].taken = value
    end
end


function M.set_free(ladder_id) -- points.set_taken(1, true)
    --print("Freeing point for ladder id:"..ladder_id)
	
	if ladder_id == hash("/ladder16") or ladder_id == hash("/ladder12") or ladder_id == hash("/ladder5") or ladder_id == hash("/ladder")then 
		M.orc_points[1].taken = false
		return 		
	end

	if ladder_id == hash("/ladder11") or ladder_id == hash("/ladder17") or ladder_id == hash("/ladder6")or ladder_id == hash("/ladder1")then 
		M.orc_points[2].taken = false
		return 		
	end

	if ladder_id == hash("/ladder14") or ladder_id == hash("/ladder15") or ladder_id == hash("/ladder7") or ladder_id == hash("/ladder2")then 
		M.orc_points[3].taken = false
		return 		
	end
	
	if ladder_id == hash("/ladder18") or ladder_id == hash("/ladder19") or ladder_id == hash("/ladder8") or ladder_id == hash("/ladder3") then 
		M.orc_points[4].taken = false
		return 		
	end

	if ladder_id == hash("/ladder10") or  ladder_id == hash("/ladder13") or ladder_id == hash("/ladder9") or ladder_id == hash("/ladder4")then 
		M.orc_points[5].taken = false
		return 		
	end

    
end

function M.get_free_point()
	local free_points = {}
	for i, point in ipairs(M.orc_points) do
		if  point.taken == false then
			table.insert(free_points, {index = i, pos = point.pos})
		end
	end
	if #free_points == 0 then 
	 return nil -- No free point available
	end
	local random_choice = math.random(#free_points) --random index from table
	return free_points[random_choice].index, free_points[random_choice].pos	
end


local savefile_path = sys.get_save_file(string.gsub(sys.get_config('project.title'), '%s+', ''), 'save_file') 
local loaded = sys.load(savefile_path)

M.sd = {
	best_score = loaded.best_score or 0
}

function M.debug_points()
	local output = ""
	for i, point in ipairs(M.orc_points) do
        output = output .. i .. ":" .. tostring(point.taken) .. " "
    end
	print(output)
end

function M.save()
	if sys.save(savefile_path, M.sd) then
		print('save successful')
	else
		print('save failed')
	end
end

return M