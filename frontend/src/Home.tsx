import { useEffect, useState } from "react";
import Sidebar from "./Sidebar";
import DonutChart from "./DonutChart";
import PushNotification from "./PushNotication";
import { useWindowSize } from "@uidotdev/usehooks";

import BlueprintsIcon from "./assets/BlueprintIcon.svg";
import SmetaIcon from "./assets/SmetaIcon.svg";

import { Gantt, Task, EventOption, StylingOption, ViewMode, DisplayOption } from 'gantt-task-react';
import "gantt-task-react/dist/index.css";

const Home = () => {

    const tasks: Task[] = [
    {
      start: new Date(2020, 1, 1),
      end: new Date(2020, 1, 2),
      name: 'Idea',
      id: 'Task 0',
      type:'task',
      progress: 45,
      isDisabled: true,
      styles: { progressColor: '#ffbb54', progressSelectedColor: '#ff9e0d' },
    },
    {
      start: new Date(2021, 1, 1),
      end: new Date(2021, 1, 2),
      name: 'Idea',
      id: 'Task 1',
      type:'task',
      progress: 45,
      isDisabled: true,
      styles: { progressColor: '#ffbb54', progressSelectedColor: '#ff9e0d' },
    },
    {
      start: new Date(2022, 1, 1),
      end: new Date(2022, 1, 2),
      name: 'Idea',
      id: 'Task 2',
      type:'task',
      progress: 45,
      isDisabled: true,
      styles: { progressColor: '#ffbb54', progressSelectedColor: '#ff9e0d' },
    }
    ];

    const [chartData, setChartData] = useState([
        { value: 0, color: "#37515f" },
        { value: 4, color: "#d2d5dd" },
    ]);

    const [userName, setUserName] = useState("");
    const [showSidebar, setSidebar] = useState(true);

    const [showNumber, setNumber] = useState(false);
    const [showEMail, setEmail] = useState(false);

    const [pushNotificationText, setPushNotificationText] = useState("");

    const toggleSidebar = () => {
        setSidebar(!showSidebar);
    };

    const pushNotification = (text : string) => {
        setPushNotificationText(text);
    }

    const checkInput = (event: React.ChangeEvent<HTMLInputElement>) => {
        const checked = event.currentTarget.checked;
        const newData = [...chartData];
        newData[0].value += checked ? 1 : -1;
        newData[1].value = 4 - newData[0].value;
        setChartData(newData);
        console.log(`1: ${chartData[0].value} 2: ${chartData[1].value}`);
    };

    const emailCopy = () => {
        const emailTextElement = document.getElementById("email");
        try {
            navigator.clipboard.writeText(emailTextElement!.innerText);
            setEmail(true);
            pushNotification("Электронная почта скопирована");
        } catch (err : unknown)
        {
            pushNotification(`Не удалось скопировать электронную почту ${err}`);
        }
    };

    const numberCopy = () => {
        const numberTextElement = document.getElementById("number");
        try {
            navigator.clipboard.writeText(numberTextElement!.innerText);
            setNumber(true);
            pushNotification("Номер телефона скопирован");
        } catch (err : unknown)
        {
            pushNotification(`Не удалось скопировать номер телефона ${err}`);
        }
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
        <main className="relative">
            {/* TODO: fix notification laggy animation and make it more stable */}
            <PushNotification text={pushNotificationText}/>

            <section className="flex flex-row w-[85%] h-[100vh] text-2xl ml-[15%]">
                <Sidebar show={showSidebar} choice={0}/>
                {/* Move to Sidebar component */}
                <button
                    onClick={toggleSidebar}
                    className={`${!showSidebar ? "ml-[-10%]" : "ml-[-50px]"} ${safeWidth <= 1100 ? "" : "hidden"} relative h-10 w-10 bg-palette-color4 rounded-2xl text-white`}
                >
                    {showSidebar ? ">" : "<"}
                </button>

                <div className="w-[80%] max-lg:w-[100vw] text-center">
                    {/* Header info */}
                    <div className="flex flex-row m-2 items-center rounded-lg border-3 p-2 border-palette-color5">
                        <section className="justify-self-start text-left ml-5">
                            {/* Recieve data from backend */}
                            <div className="text-blue-950 font-bold text-3xl mb-5">
                                Строительство дома
                            </div>

                            <div className="mb-5">
                                {userName === "unknown"
                                    ? "Пользователь не вошел!"
                                    : `Заказчик: ${userName}`}
                            </div>

                            {/* Recieve data from backend */}
                            <div>Выполняется</div>
                        </section>
                        
                        <section className="ml-auto self-end transition-all">
                            <DonutChart
                                data={chartData}
                                width={200}
                                height={200}
                                innerRadius={90}
                                outerRadius={100}
                                
                            />
                        </section>
                    </div>

                    {/* Tasks panel */}
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
                                    className="mr-2"
                                />
                                <label htmlFor="1">Заливка фундамента</label>
                            </li>
                            <li>
                                <input
                                    onChange={checkInput}
                                    type="checkbox"
                                    id="2"
                                    className="mr-2"
                                />
                                <label htmlFor="2">Возведение стен</label>
                            </li>
                            <li>
                                <input
                                    onChange={checkInput}
                                    type="checkbox"
                                    id="3"
                                    className="mr-2"
                                />
                                <label htmlFor="3">Монтаж кровли</label>
                            </li>
                            <li>
                                <input
                                    onChange={checkInput}
                                    type="checkbox"
                                    id="4"
                                    className="mr-2"
                                />
                                <label htmlFor="4">Устройство террасы</label>
                            </li>
                        </ul>
                    </div>
                    
                    <div className="w-[98.5%] rounded-lg max-w-full overflow-x-auto border-3 m-2 border-palette-color5 flex items-center">
                        <Gantt tasks={tasks} />
                    </div>

                    {/* Contact panel */}
                    <div className="flex flex-row m-2 items-center rounded-lg border-3 border-palette-color5 text-left p-4">
                        <div className="text-blue-950 font-bold">
                            Связь с клиентом:
                        </div>
                        <section className="flex flex-row ml-auto self-end items-center space-x-2">
                            <div id="number" className={`absolute ${showNumber ? "" : "pointer-events-none"} select-none text-center justify-self-center transition-opacity duration-300 ${showNumber ? "opacity-100" : "opacity-0"}`}>
                                +78005353535
                            </div>
                            <input
                                data-copy-to-clipboard-target="number"
                                onClick={() => numberCopy()}
                                type="button"
                                value="Показать номер"
                                className={`border-2 border-blue-950 p-2 rounded-xl transition-opacity duration-300 ${showNumber ? "opacity-0" : "opacity-100"}`}
                            />
                        </section>
                        <section className="flex flex-row ml-auto self-end items-center space-x-2 mt-2 mr-15">
                            <div id="email" className={`absolute ${showEMail ? "" : "pointer-events-none"} select-none text-center justify-self-center transition-opacity duration-300 ${showEMail ? "opacity-100" : "opacity-0"}`}>
                                aelilkidae@gmail.com
                            </div>
                            <input
                                data-copy-to-clipboard-target="email"
                                onClick={() => emailCopy()}
                                type="button"
                                value="Показать почту"
                                className={`border-2 border-blue-950 p-2 rounded-xl transition-opacity duration-300 ${showEMail ? "opacity-0" : "opacity-100"}`}
                            />
                        </section>
                    </div>

                    {/* Empty div */}
                    <div className="w-full h-5"></div>
                </div>

		        <div className="flex flex-col w-[20%] h-screen p-2 gap-2">
		            <a href="/estimates" className="flex flex-col w-full h-1/3 border-3 border-palette-color5 justify-center items-center rounded-lg">
		                <div className="flex items-center justify-center w-[60%] h-[60%] bg-palette-color2 rounded-full">
		            	    <img src={SmetaIcon} className="max-w-[55%] max-h-[55%]"/>  
		            	</div>
		        	<div className="mt-5 font-semibold font-[#1d3266]">Создать смету</div>
		            </a>
		            <a href="/blueprints" className="flex flex-col w-full h-1/3 border-3 border-palette-color5 justify-center items-center rounded-lg">
		                <div className="flex items-center justify-center w-[60%] h-[60%] bg-palette-color2 rounded-full">
		            	    <img src={BlueprintsIcon} className="max-w-[55%] max-h-[55%]"/>  
		            	</div>
		        	<div className="mt-5 font-semibold font-[#1d3266]">Открыть чертежи</div>
		            </a>
		            <a href="/projects" className="flex flex-col w-full h-1/3 border-3 border-palette-color5 justify-center items-center rounded-lg">
		                <span className="material-symbols-outlined-huge">settings</span>
		        	<div className="mt-5 font-semibold font-[#1d3266]">Открыть проекты</div>
		            </a>
		        </div>
            </section>
        </main>
    );
};

export default Home;
