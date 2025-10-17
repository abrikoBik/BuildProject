import { useEffect, useState } from "react";
import Sidebar from "./Sidebar";
import PushNotification from "./PushNotication";
import { useWindowSize } from "@uidotdev/usehooks";
import Calendar from "react-calendar";
import axios from "axios";

import { ProjectRequest } from "./types/projects";
import { UserResponse } from "./types/user";
import { TaskGroup, TaskRequest, Task } from "./types/home";

import GanttChart from "./Gantt";
import { Chart } from "react-google-charts";
import Header from "./Header";

type ValuePiece = Date | null;

type Value = ValuePiece | [ValuePiece, ValuePiece];

const Home = () => {

    {/* TODO: remove role from localstorage */}
    const role = localStorage.getItem("role");

    const data = [
        ["xyi", "xyi"],
        ["Work", 9],
        ["Eat", 2],
        ["Commute", 2],
        ["Watch TV", 2],
        ["Sleep", 7],
    ];

    const options = {
        pieHole: 0.6,
        legend: 'none',
        pieSliceText: 'none'
    };
    
    const [selectedProject, setSelectedProject] = useState<ProjectRequest>()

    const [date, setDate] = useState<Value>(new Date());

    const [user, setUser] = useState<UserResponse>();
    const [showSidebar, setSidebar] = useState(true);

    const [showNumber, setNumber] = useState(false);
    const [showEMail, setEmail] = useState(false);

    const [addTaskGroupInput, setAddTaskGroupInput] = useState(false);
    const [taskGroupsList, setTaskGroupsList] = useState<TaskGroup[]>();

    const [taskData, setTaskData] = useState<TaskRequest>({
        startTime: 0,
        endTime: 0,
        name: "",
    });

    const [tasksList, setTasksList] = useState<Task[]>()

    const [pushNotificationText, setPushNotificationText] = useState("");

    const toggleSidebar = () => {
        setSidebar(!showSidebar);
    };

    const pushNotification = (text : string) => {
        setPushNotificationText(text);
    }

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

    const taskGroupAdd = async (name : string) => {
        try {
            const taskGroup = await axios.post(
                `api/projects/${localStorage.getItem("selectedProjectID")}/stageGroups`,
                {"name": name},
                {
                    headers: {
                        "Content-Type": "application/json"
                    }
                }
            );

            console.log(taskGroup.status, taskGroup.data);
            listTaskGroups();

        } catch (e : unknown) {
            console.log(e);
        }
    }

    const listTaskGroups = async () => {
        try {
            const taskGroup = await axios.get(
                `api/projects/${localStorage.getItem("selectedProjectID")}/stageGroups`
            );

            setTaskGroupsList(taskGroup.data);
        } catch (e : unknown) {
            console.log(e);
        }
    }

    const listTasks = async () => {
        try {
            const tasks = await axios.get<Task[]>(
                `api/projects/${localStorage.getItem("selectedProjectID")}/stageGroups/${localStorage.getItem("selectedStageGroupID")}/stages`
            )

            setTasksList(tasks.data);
        } catch (e : unknown) {
            console.log(e);
        }
    }

    const addTask = async () => {
        try {
            const tasks = await axios.post(
                `api/projects/${localStorage.getItem("selectedProjectID")}/stageGroups/${localStorage.getItem("selectedStageGroupID")}/stages`,
                taskData,
                {
                    headers: {
                        "Content-Type": "application/json"
                    }
                }
            )

            console.log(tasks.data);
            console.log(taskData);
            listTasks();
        } catch (e : unknown) {
            console.log(e);
        }
    }

    const handleTaskInput = (e : React.ChangeEvent<HTMLInputElement>) => {
        if(e.target.id === "name")
        {
            setTaskData({
                ...taskData, 
                name: e.target.value,
            })
        } else if (e.target.id === "startDate") {
            setTaskData({
                ...taskData, 
                startTime: Math.floor((new Date(e.target.value)).getTime() / 1000)
            })
        } else {
            setTaskData({
                ...taskData,
                endTime: Math.floor((new Date(e.target.value)).getTime() / 1000)
            })
        }
    }

    const deleteTask = async (id : number) => {
        try {
            await axios.delete(
                `api/projects/${localStorage.getItem("selectedProjectID")}/stageGroups/${localStorage.getItem("selectedStageGroupID")}/stages/${id}`
            )

            listTasks();
        } catch (e : unknown) {
            console.log(e);
        }
    }

    const deleteTaskGroup = async (id : number) => {
        try {
            const taskGroup = await axios.delete(
                `api/projects/${localStorage.getItem("selectedProjectID")}/stageGroups/${id}/`
            )

            console.log(taskGroup.status);
            listTaskGroups();
            listTasks();
        } catch (e : unknown) {
            console.log(e);
        }
    }

    const unixDateStampToDate = (unixTimeStamp : number) => {
        const date = new Date(unixTimeStamp * 1000);
        return `${date.getDate()}/${date.getMonth()+1}/${date.getFullYear()}`;
    }

    const { width } = useWindowSize();
    const safeWidth = width ?? 1200;

    useEffect(() => {
        setSidebar(safeWidth > 1100);
    }, [safeWidth]);

    useEffect(() => {
        getProject();
        listTaskGroups();
        listTasks();
        console.log(role);
    }, [])

    return (
        <main className="relative">
            <Header />
            <PushNotification text={pushNotificationText}/>
            <section className="flex flex-row w-[85vw]  h-[100vh] text-2xl ml-[15%]">
                <Sidebar show={showSidebar} choice={0}/>
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
                    <div className="flex flex-row m-2 rounded-lg border-3 p-2 border-palette-color5">
                        <section className="justify-self-start text-left ml-5">
                            <div className="text-blue-950 font-bold text-3xl mb-5">
                                {selectedProject?.name}
                            </div>

                            <div className="mb-5">
                                {user?.name}
                            </div>

                            <div>Выполняется</div>
                        </section>
                        
                        <section className="ml-auto self-end transition-all">
                            <Chart
                              chartType="PieChart"
                              data={data}
                              options={options}
                              width={"300px"}
                              height={"300px"}
                            />
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
                                        <input autoFocus type="text" id="groupInput" className="w-full !text-black !border-0" onKeyDown={(e) => {
                                            if(e.key === "Enter")
                                            {
                                                setAddTaskGroupInput(false);
                                                if(e.currentTarget.value != "")
                                                    taskGroupAdd(e.currentTarget.value);
                                            } else if (e.key === "Escape") {
                                                setAddTaskGroupInput(false);
                                            }
                                        }} />
                                    </li>
                                    :
                                    <></>
                                }

                                { taskGroupsList?.map((item, index) => (
                                    <li 
                                        key={index}
                                        onClick={() => {localStorage.setItem("selectedStageGroupID", item.id.toString()); listTaskGroups(); listTasks();}}
                                        className={`w-full min-h-10 ${item.id === parseInt(localStorage.getItem("selectedStageGroupID")!) ? "bg-yellow-100" : ""} items-center flex px-3 border-3 border-palette-color5 rounded-lg justify-between hover:shadow-[0px_0px_5px_#37515fff]`}
                                    >
                                        <div>{item.name}</div>
                                        <span onClick={() => deleteTaskGroup(item.id)} className="material-symbols-outlined hover:text-red-700 select-none hover:cursor-pointer">
                                            close
                                        </span>
                                    </li>
                                ))
                                }
                            </ul>
                            { role === "client" ?
                                <></>
                                :
                                <button onClick={() => {
                                        setAddTaskGroupInput(!addTaskGroupInput);
                                        if((document.getElementById("groupInput") as HTMLInputElement).value)
                                            taskGroupAdd((document.getElementById("groupInput") as HTMLInputElement).value);
                                    }} 
                                    className="flex w-auto px-2 items-center  active:text-white rounded border-3 self-center bg-palette-color4 active:!bg-[#232c33] text-white border-palette-color4 hover:!bg-white hover:text-black hover:cursor-pointer"
                                >
                                Добавить
                                </button>
                            }
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

                            <ul className="flex flex-col h-[30vh] self-start gap-3 ml-5 overflow-scroll">
                                { tasksList?.map((item, index) => (
                                        <li key={index} className="flex flex-row justify-between pr-5">
                                            <section>
                                                { role === "client" ?
                                                <input type="checkbox" id="checked" checked={item.status} disabled className="mr-2"/>
                                                :
                                                <input
                                                    type="checkbox"
                                                    id="check"
                                                    checked={item.status}
                                                    onChange={async (e : React.ChangeEvent<HTMLInputElement>) => 
                                                        {
                                                            await axios.patch(
                                                                `api/projects/${localStorage.getItem("selectedProjectID")}/stageGroups/${localStorage.getItem("selectedStageGroupID")}/stages/${item.id}`,
                                                                {...item, status: !item.status}
                                                            )

                                                            listTasks();
                                                            console.log(tasksList);
                                                        }}
                                                    className="mr-2"
                                                />
                                                }
                                                <label htmlFor="check">{item.name}</label>
                                            </section>
                                            <div className="text-gray-500">
                                                {`${unixDateStampToDate(item.startTime)} -> ${unixDateStampToDate(item.endTime)}`}
                                                <span onClick={() => deleteTask(item.id)}className="material-symbols-outlined hover:text-red-700 select-none">
                                                    close
                                                </span>
                                            </div>
                                            
                                        </li>
                                    ))
                                }
                            </ul>
                            { role === "client" 
                                ?
                                <></>
                                :
                                <section className="flex flex-col gap-3 ml-5">
                                    <div className="w-full">
                                        <input 
                                            onKeyDown={(e : React.KeyboardEvent) => {
                                                if(e.key === "Enter") {
                                                    addTask();
                                                }
                                            }}
                                            type="text"
                                            id="name"
                                            onChange={handleTaskInput}
                                            className="w-[45%]"
                                        />
                                        <input
                                            type="date"
                                            id="startDate"
                                            onChange={handleTaskInput}
                                        />
                                        <input
                                            type="date"
                                            id="endDate"
                                            onChange={handleTaskInput}
                                        />
                                    </div>
                                    <div>
                                        <button 
                                            onClick={() => {
                                                addTask();
                                            }}
                                            className="w-full min-h-10 bg-palette-color4 text-white rounded-xl  hover:!bg-white hover:text-black border-3 border-palette-color4 hover:cursor-pointer active:!bg-[#232c33] active:text-white"
                                        >
                                        Добавить задачу
                                        </button>
                                    </div>
                                </section>
                            }
                        </section>
                    </div>
                    
                    <div className="flex-row w-[98.5%] rounded-lg max-w-full overflow-x-auto border-3 m-2 border-palette-color5 flex items-center">
                        <div className="flex flex-row w-[75%]">
                            <GanttChart />
                        </div>
                        <Calendar onChange={setDate} value={date}/>
                    </div>
                    { role === "client" ?
                        <></>
                        :
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
                    }

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
