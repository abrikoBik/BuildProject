import { useEffect, useState } from "react";
import Sidebar from "./Sidebar";
import DonutChart from "./DonutChart";

const Home = () => {
    const [chartData, setChartData] = useState([
        { value: 0, color: "#37515f" },
        { value: 4, color: "#d2d5dd" },
    ]);

    const [userName, setUserName] = useState("");
    const [showSidebar, setSidebar] = useState(true);

    const toggleSidebar = () => {
        setSidebar(!showSidebar);
    };

    useEffect(() => {
        let newName: string;
        try {
            newName = JSON.parse(localStorage.getItem("user")!).name;
            setUserName(newName);
        } catch {
            setUserName("unknown");
            console.warn("USER DONT LOGIN!");
        }
    }, []);

    const checkInput = (event: React.ChangeEvent<HTMLInputElement>) => {
        const checked = event.currentTarget.checked;
        const newData = [...chartData];
        newData[0].value += checked ? 1 : -1;
        newData[1].value = 4 - newData[0].value;
        setChartData(newData);
        console.log(`1: ${chartData[0].value} 2: ${chartData[1].value}`);
    };

    return (
        <main>
            <section className="flex flex-row w-[100vw] h-[100vh] text-2xl">
                <Sidebar show={showSidebar} />
                <button
                    onClick={toggleSidebar}
                    className={`${!showSidebar ? "" : "hidden"}`}
                >
                    {">"}
                </button>

                <div className="w-[60vw] max-md:w-[100vw] text-center">
                    <div className="flex flex-row m-2 items-center rounded-lg border-2 p-2 border-palette-color5">
                        <section className="justify-self-start text-left ml-5">
                            {/* Брать данные с бэка */}
                            <div className="text-blue-950 font-bold text-3xl mb-5">
                                Строительство дома
                            </div>

                            <div className="mb-5">
                                {userName === "unknown"
                                    ? "Пользователь не вошел!"
                                    : `Заказчик: ${userName}`}
                            </div>

                            {/* Брать данные с бэка */}
                            <div>Выполняется</div>
                        </section>
                        
                        <section className="ml-auto self-end">
                            <DonutChart
                                data={chartData}
                                width={200}
                                height={200}
                                innerRadius={80}
                                outerRadius={100}
                                
                            />
                        </section>
                    </div>

                    <div className="flex flex-row m-2 items-center rounded-lg border-2 p-2 border-palette-color5 text-left p-4">
                        <div className="text-blue-950 font-bold">
                            Связь с клиентом:
                        </div>
                        <input type="button" value="Показать номер" className="border-2 border-blue-950 ml-auto self-end p-1 rounded-xl px-2"/>
                    </div>

                    <div className="flex flex-col m-2 items-center rounded-lg border-2 p-2 border-palette-color5 text-left">
                        <section className="flex flex-row w-[100%]">
                            <div className="text-blue-950 font-bold text-3xl ml-5 self-start">
                                Задачи
                            </div>

                            <div className="text-blue-950 font-bold text-3xl ml-auto mr-5 self-end">
                                Дата
                            </div>
                        </section>

                        <ul className="flex flex-col self-start gap-3 ml-5">
                            <li>
                                <input
                                    onChange={checkInput}
                                    type="checkbox"
                                    id="1"
                                />
                                <label htmlFor="1">Заливка фундамента</label>
                            </li>
                            <li>
                                <input
                                    onChange={checkInput}
                                    type="checkbox"
                                    id="2"
                                />
                                <label htmlFor="2">Возведение стен</label>
                            </li>
                            <li>
                                <input
                                    onChange={checkInput}
                                    type="checkbox"
                                    id="3"
                                />
                                <label htmlFor="3">Монтаж кровли</label>
                            </li>
                            <li>
                                <input
                                    onChange={checkInput}
                                    type="checkbox"
                                    id="4"
                                />
                                <label htmlFor="4">Устройство террасы</label>
                            </li>
                        </ul>
                    </div>
                </div>
            </section>
        </main>
    );
};

export default Home;
