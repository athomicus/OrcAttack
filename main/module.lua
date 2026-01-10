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

function M.free_ladder_for_next_orc(index)
	M.orc_points[index].taken = false
end



function M.get_free_point()
	local free_points = {}  -- 1. Pusta tablica na wolne punkty
	for i, point in ipairs(M.orc_points) do -- 2. Przeglądaj WSZYSTKIE punkty
		if  point.taken == false then -- 3. Jeśli punkt WOLNY (nie zajęty)
			table.insert(free_points, {index = i, pos = point.pos}) -- 4. Dodaj do listy: numer + pozycja
			--pprint(free_points[1].index)
		end
	end
	if #free_points == 0 then 
	 return nil -- No free point available
	end
	local random_choice = math.random(#free_points) --random index from table
	return free_points[random_choice].index, free_points[random_choice].pos	
	-- 8. Zwróć: INDEX - czyli 1 lub 2 lub 3lub 4 lub 5 drabina i POZYCJĘ -x,y,z losowego wolnego punktu
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