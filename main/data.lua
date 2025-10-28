local M = {}


local savefile_path = sys.get_save_file(string.gsub(sys.get_config('project.title'), '%s+', ''), 'save_file') 
local loaded = sys.load(savefile_path)

M.sd = {
	best_score = loaded.best_score or 0
}

function M.save()
	if sys.save(savefile_path, M.sd) then
		print('save successful')
	else
		print('save failed')
	end
end

return M