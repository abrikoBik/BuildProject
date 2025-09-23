const ManagerPanel = () => {
    return (
        <main>
            <h2 className="font-bold text-2xl m-1">Панель редактирования ролей</h2>
            <div className="ml-2 mt-2">
                <label htmlFor="roles" className="mr-2 !text-black">Выберите роль</label>
                <select name="roles" id="roles" className="bg-gray-300 p-1 border-2">
                    <option value="worker">worker</option>
                    <option value="client">client</option>
                </select>
            </div>            
        </main>
    );
}

export default ManagerPanel;