import { useEffect, useState } from "react";
import { Sidebar } from "./Sidebar";
import { DonutChart } from "./DonutChart"

export const Home = () => {
    const [chartData, setChartData] = useState([
        { value: 0, color: '#37515f' },
        { value: 4, color: '#d2d5dd'},
    ]);

    const [userName, setUserName] = useState('');
    const [showSidebar, setSidebar] = useState(true);

    const toggleSidebar = () => {
        setSidebar(!showSidebar);
    }

    useEffect(() => {
        let newName : string;
        try {
            newName = JSON.parse(localStorage.getItem('user')!).name;
            setUserName(newName);
        } catch {
            setUserName('unknown');
            console.warn("USER DONT LOGIN!");
        }
    },[])

    const checkInput = (event: React.ChangeEvent<HTMLInputElement>) => {
        const checked = event.currentTarget.checked;
        const newData = [...chartData];
        newData[0].value += checked ? 1 : -1;
        newData[1].value = 4 - newData[0].value;
        setChartData(newData);
        console.log(`1: ${chartData[0].value} 2: ${chartData[1].value}`);
    }

    return (
        <main>
            <section className="flex flex-row w-[100vw] h-[100vh] text-2xl">
                <Sidebar show={showSidebar}/>
                <button onClick={toggleSidebar} className={`${!showSidebar ? '' :'hidden'}`}>
                    {'>'}
                </button>

                <div className="w-[85vw] max-md:w-[100vw] text-center">
                    <div className="flex flex-row items-center justify-center m-2 rounded-lg border-2 p-2 border-palette-color5">
                        <section>
                            {/* Брать данные с бэка */}
                            <div className="text-blue-950">
                                Строительство дома
                            </div>

                            <div>
                                { userName === 'unknown' ? 'Пользователь не вошел!' : `Заказчик: ${userName}` }
                            </div>
                            
                            {/* Брать данные с бэка */}
                            <div>
                                Выполняется
                            </div>
                        </section>

                        <DonutChart 
                            data={chartData}
                            width={200}
                            height={200}
                            innerRadius={80}
                            outerRadius={100}
                        />
                    </div>

                    <ul>
                        <li><input onChange={checkInput} type="checkbox" id="1"/><label htmlFor="1">Заливка фундамента</label></li>
                        <li><input onChange={checkInput} type="checkbox" id="2"/><label htmlFor="2">Возведение стен</label></li>
                        <li><input onChange={checkInput} type="checkbox" id="3"/><label htmlFor="3">Монтаж кровли</label></li>
                        <li><input onChange={checkInput} type="checkbox" id="4"/><label htmlFor="4">Устройство террасы</label></li>
                    </ul>
                </div>
            </section>
        </main>
    );
};
