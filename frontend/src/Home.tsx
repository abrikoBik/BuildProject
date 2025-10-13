import { useEffect, useState } from "react";
import Sidebar from "./Sidebar";
import DonutChart from "./DonutChart";
import PushNotification from "./PushNotication";
import { useWindowSize } from "@uidotdev/usehooks";
import Calendar from "react-calendar";
import axios from "axios";

import { ProjectRequest } from "./types/projects";
import { UserResponse } from "./types/user";

import GanttChart from "./Gantt";
import Header from "./Header";

type ValuePiece = Date | null;

type Value = ValuePiece | [ValuePiece, ValuePiece];

const Home = () => {

    const [chartData, setChartData] = useState([
        { value: 0, color: "#37515f" },
        { value: 6, color: "#d2d5dd" },
    ]);

    const [selectedProject, setSelectedProject] = useState<ProjectRequest>()

    const [date, setDate] = useState<Value>(new Date());

    const [user, setUser] = useState<UserResponse>();
    const [showSidebar, setSidebar] = useState(true);

    const [showNumber, setNumber] = useState(false);
    const [showEMail, setEmail] = useState(false);

    const [addTaskGroupInput, setAddTaskGroupInput] = useState(false);

    const [pushNotificationText, setPushNotificationText] = useState("");

    const toggleSidebar = () => {
        setSidebar(!showSidebar);
    };

    const pushNotification = (text : string) => {
        setPushNotificationText(text);
    }

    const checkInput = (event: React.ChangeEvent<HTMLInputElement>) => {
        const chartMaxValue = chartData[0].value + chartData[1].value;
        const checked = event.currentTarget.checked;
        const newData = [...chartData];
        newData[0].value += checked ? 1 : -1;
        newData[1].value = chartMaxValue - newData[0].value;
        setChartData(newData);
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

    const getProject = async () => {
        try {
            const project = await axios.get<ProjectRequest>(
                `/api/projects/${localStorage.getItem("selectedProjectID")}`
            );

            if(project.status != 200)
            {
                console.log(project.status);
                return;
            }

            setSelectedProject(project.data);

            const user = await axios.get<UserResponse>(
                `/api/users/${project.data.userID}`
            );

            setUser(user.data);

        } catch (e : unknown) {
            console.error(e);
        }
    }

    const { width } = useWindowSize();
    const safeWidth = width ?? 1200;

    useEffect(() => {
        setSidebar(safeWidth > 1100);
    }, [safeWidth]);

    useEffect(() => {
        getProject();
    }, [])

    return (
        <main className="relative">
            <Header />
            <PushNotification text={pushNotificationText}/>
            <section className="flex flex-row w-[85vw]  h-[100vh] text-2xl ml-[15%]">
                <Sidebar show={showSidebar} choice={0}/>
                {/* Move to Sidebar component */}
                <button
                    onClick={toggleSidebar}
                    className={`${!showSidebar ? "ml-[-10%]" : "ml-[-50px]"} ${safeWidth <= 1100 ? "" : "hidden"} relative h-10 w-10 bg-palette-color4 rounded-2xl text-white`}
                >
                    {showSidebar ? ">" : "<"}
                </button>
                { localStorage.getItem("selectedProjectID")  ?
                <>
                <div className="w-[85vw] mt-20 max-lg:w-[100vw] text-center">
                    {/* Header info */}
                    <div className="flex flex-row m-2 items-center rounded-lg border-3 p-2 border-palette-color5">
                        <section className="justify-self-start text-left ml-5">
                            {/* Recieve data from backend */}
                            <div className="text-blue-950 font-bold text-3xl mb-5">
                                {selectedProject?.name}
                            </div>

                            <div className="mb-5">
                                {user?.name}
                            </div>

                            {/* Recieve data from backend */}
                            <div>Выполняется</div>
                        </section>
                        
                        <section className="ml-auto self-end transition-all">
                            google charts
                        </section>
                    </div>

                    {/* Tasks panel */}
                    <div className="flex flex-row m-2 rounded-lg border-3 p-2 border-palette-color5 text-left">
                        <section className="flex flex-col w-1/2 h-full">
                            <h2 className="text-blue-950 font-bold text-3xl ml-5 mb-3 self-start">
                                Группы задач
                            </h2>

                            <ul className="flex flex-col h-70 overflow-scroll gap-3 mb-5">
                                {
                                    addTaskGroupInput ?
                                    <li className="w-full h-10 items-center flex px-3 border-3 border-palette-color5 rounded-lg">
                                        <input autoFocus type="text" id="groupInput" className="w-full !text-black" onKeyDown={(e) => {
                                            if(e.key === "Enter")
                                            {
                                                setAddTaskGroupInput(false);
                                            }
                                        }} />
                                    </li>
                                    :
                                    <></>
                                }
                                <li className="w-full h-10 items-center flex px-3 border-3 border-palette-color5 rounded-lg">Заливка фундамента 1</li>
                            </ul>

                            <button onClick={() => {
                                    setAddTaskGroupInput(!addTaskGroupInput);
                                }} 
                                className="flex w-auto px-2 items-center rounded border-3 self-center"
                            >
                            Добавить
                            </button>
                        </section>

                        <section className="w-1/2 h-full">
                            <section className="flex flex-row w-[100%]">
                                <h2 className="text-blue-950 font-bold text-3xl ml-5 self-start">
                                    Задачи
                                </h2>

                                <h2 className="text-blue-950 font-bold text-3xl ml-auto mr-5 self-end">
                                    Дата
                                </h2>
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
                                <li>
                                    <input
                                        onChange={checkInput}
                                        type="checkbox"
                                        id="5"
                                        className="mr-2"
                                    />
                                    <label htmlFor="5">Устройство террасы</label>
                                </li>
                                <li>
                                    <input
                                        onChange={checkInput}
                                        type="checkbox"
                                        id="6"
                                        className="mr-2"
                                    />
                                    <label htmlFor="6">Устройство террасы</label>
                                </li>
                            </ul>
                        </section>
                    </div>
                    
                    <div className="flex-row w-[98.5%] rounded-lg max-w-full overflow-x-auto border-3 m-2 border-palette-color5 flex items-center">
                        <div className="flex flex-row w-[75%]">
                            <GanttChart />
                        </div>
                        <Calendar onChange={setDate} value={date}/>
                    </div>

                    {/* Contact panel */}
                    <div className="flex flex-row m-2 items-center rounded-lg border-3 border-palette-color5 text-left p-4">
                        <div className="text-blue-950 font-bold">
                            Связь с клиентом:
                        </div>
                        <section className="ml-[30%] flex flex-row gap-20">
                            <section className="flex flex-row ml-auto items-center">
                                <div id="number" className={`absolute ${showNumber ? "" : "pointer-events-none"} select-none text-center justify-self-center transition-opacity duration-300 ${showNumber ? "opacity-100" : "opacity-0"}`}>
                                    {user?.phoneNumber}
                                </div>
                                <input
                                    data-copy-to-clipboard-target="number"
                                    onClick={() => numberCopy()}
                                    type="button"
                                    value="Показать номер"
                                    className={`border-2 border-blue-950 p-2 rounded-xl transition-opacity duration-300 ${showNumber ? "opacity-0" : "opacity-100"}`}
                                />
                            </section>
                            <section className="flex flex-row items-center overflow-x-scroll">
                                <div id="email" className={`absolute ${showEMail ? "" : "pointer-events-none"} select-none text-center justify-self-center transition-opacity duration-300 ${showEMail ? "opacity-100" : "opacity-0"}`}>
                                    {user?.email}
                                </div>
                                <input
                                    data-copy-to-clipboard-target="email"
                                    onClick={() => emailCopy()}
                                    type="button"
                                    value="Показать почту"
                                    className={`border-2 border-blue-950 p-2 rounded-xl transition-opacity duration-300 ${showEMail ? "opacity-0" : "opacity-100"}`}
                                />
                            </section>
                        </section>
                    </div>

                    {/* Empty div */}
                    <div className="w-full h-5"></div>
                </div>
                </>
                :
                <div className="flex w-full h-full justify-center items-center">Проект не выбран!</div>
                }
            </section>
        </main>
    );
};

export default Home;
