import { useEffect, useState } from "react";
import Sidebar from "./Sidebar";
import DonutChart from "./DonutChart";
import { useWindowSize } from "@uidotdev/usehooks";

const Home = () => {
    const [chartData, setChartData] = useState([
        { value: 0, color: "#37515f" },
        { value: 4, color: "#d2d5dd" },
    ]);

    const [userName, setUserName] = useState("");
    const [showSidebar, setSidebar] = useState(true);

    const [showNumber, setNumber] = useState(false);
    const [showEMail, setEmail] = useState(false);

    const toggleSidebar = () => {
        setSidebar(!showSidebar);
    };

    const checkInput = (event: React.ChangeEvent<HTMLInputElement>) => {
        const checked = event.currentTarget.checked;
        const newData = [...chartData];
        newData[0].value += checked ? 1 : -1;
        newData[1].value = 4 - newData[0].value;
        setChartData(newData);
        console.log(`1: ${chartData[0].value} 2: ${chartData[1].value}`);
    };

    const { width } = useWindowSize();
    const safeWidth = width ?? 1200;

    useEffect(() => {
        let newName: string;
        try {
            newName = JSON.parse(localStorage.getItem("user")!).name;
            setUserName(newName);
        } catch {
            setUserName("unknown");
            console.warn("USER DONT LOGIN!");
        }

        // Adaptive 

        setSidebar(safeWidth > 1100);

    }, [safeWidth]);

    return (
        <main>
            <section className="flex flex-row w-[80vw] h-[100vh] text-2xl ml-[17vw]">
                <Sidebar show={showSidebar} choice={0}/>
                <button
                    onClick={toggleSidebar}
                    className={`${!showSidebar ? "" : "ml-[-50px]"} ${safeWidth <= 1100 ? "" : "hidden"} relative h-10 w-10 bg-palette-color4 rounded-2xl text-white`}
                >
                    {showSidebar ? "<" : ">"}
                </button>

                <div className="w-[60vw] max-lg:w-[100vw] text-center">
                    <div className="flex flex-row m-2 items-center rounded-lg border-3 p-2 border-palette-color5">
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

                    <div className="flex flex-row m-2 items-center rounded-lg border-3 p-2 border-palette-color5 text-left p-4">
                        <div className="text-blue-950 font-bold">
                            Связь с клиентом:
                        </div>
                        <section className="flex flex-row ml-auto self-end items-center space-x-2">
                            <div className={`absolute z-50 pointer-events-none text-center justify-self-center transition-opacity duration-300 ${showNumber ? "opacity-100" : "opacity-0"}`}>
                                +78005353535
                            </div>
                            <input
                                onClick={() => setNumber(!showNumber)}
                                type="button"
                                value="Показать номер"
                                className={`border-2 border-blue-950 p-2 rounded-xl transition-opacity duration-300 ${showNumber ? "opacity-0" : "opacity-100"}`}
                            />
                        </section>
                        <section className="flex flex-row ml-auto self-end items-center space-x-2 mt-2 mr-15">
                            <div className={`absolute pointer-events-none text-center justify-self-center transition-opacity duration-300 ${showEMail ? "opacity-100" : "opacity-0"}`}>
                                aelilkidae@gmail.com
                            </div>
                            <input
                                onClick={() => setEmail(!showEMail)}
                                type="button"
                                value="Показать почту"
                                className={`border-2 border-blue-950 p-2 rounded-xl transition-opacity duration-300 ${showEMail ? "opacity-0" : "opacity-100"}`}
                            />
                        </section>
                    </div>

                    <div className="flex flex-col m-2 items-center rounded-lg border-3 p-2 border-palette-color5 text-left">
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
