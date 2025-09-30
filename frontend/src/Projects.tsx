import Sidebar from './Sidebar';
import Header from './Header';
import { useState, useEffect } from 'react';
import Calendar from 'react-calendar';
import 'react-calendar/dist/Calendar.css';
import { Value, ProjectRequest } from './types/projects'
import axios from 'axios';

const Projects = () => {
    const [userRequest, setUserRequest] = useState<ProjectRequest>({
        name: "",
        description: "",
        userID: JSON.parse(localStorage.getItem("user") as string).userID,
        startTime: 1,
        endTime: 1,
        latitude: 0,
        longitude: 0,
        id: 0
    });

    const [showPopUp, setShowPopUp] = useState(false);
    
    const [projectsList, setProjectsList] = useState<ProjectRequest[]>();

    const [maximizedProjects, setProjectMaximize] = useState<boolean[]>();

    const [startUpDate, setStartUpDate] = useState<Value>(new Date());
    const [endDate, setEndDate] = useState<Value>(new Date());

    const handleInput = (e : React.ChangeEvent<HTMLInputElement> | React.ChangeEvent<HTMLTextAreaElement>) => {
        setUserRequest({
            ...userRequest,
            [e.target.id]: e.target.value,
            userID: parseInt(localStorage.getItem("userID")!),
            startTime: Math.floor((startUpDate as Date).getTime() / 1000),
            endTime:   Math.floor((endDate as Date).getTime() / 1000)
        })
    }

    const createProject = async () => {
        try {

            console.log(userRequest);

            const { data, status } = await axios.post(
                '/api/projects',
                userRequest,
                {
                    headers: {
                        "Content-Type": "application/json"
                    }
                }
            );

            if(status != 200)
            {
                alert(status);
                return;
            }

            console.log(data);

            setShowPopUp(false);

            listProjects();
        } catch (err : unknown) {
            console.log(err);
        }
    }

    const deleteProject = async (id : number) => {
        try {
            const project = await axios.delete(
                `/api/projects/${id}/`
            );

            if(project.status != 200) {
                console.error(project.status);
                return;
            }
            
            listProjects();
        } catch (err : unknown) {
            console.error(err);
        }
    }

    const listProjects = async () => {
        try {
            const { data, status } = await axios.get(
                `/api/users/${localStorage.getItem("userID")}/projects`
            );

            if(status != 200) {
                alert(status);
                return;
            }

            console.log(data);
            
            setProjectsList(data);
            setProjectMaximize(Array.from({length: data.length}, () => false))

        } catch (err : unknown) {
            console.log(err);
        }
    }

    const unixDateStampToDate = (unixTimeStamp : number) => {
        const date = new Date(unixTimeStamp * 1000);
        return date;
    }

    const selectProject = (projectID: number) => {
        localStorage.setItem("selectedProjectID", projectID.toString());
        window.location.href = '/';
    }

    useEffect(() => {
        listProjects();
    }, [])

    return (
        <div className="flex flex-col h-screen">
            <Header />

            { showPopUp ? 
                <section className="fixed left-0 top-0 w-screen h-screen bg-[rgba(0,0,0,0.4)] z-150">
                    <div className="flex flex-col relative left-1/2 top-1/2 -translate-1/2 w-[60%] h-[80%] rounded-xl border-2 border-palette-color4 bg-palette-color1 overflow-y-scroll">
                        <button onClick={() => setShowPopUp(false)} type="button" className="hover:cursor-pointer hover:text-red-700 w-7 h-7 flex items-center justify-center self-end">
                            <span className="material-symbols-outlined">
                                close
                            </span>
                        </button>

                        <form action={createProject} className="flex flex-col w-full h-full p-5 gap-5">
                            <div className="flex flex-row gap-2">
                                <label htmlFor="name" className="!text-black">Название проекта</label>
                                <input 
                                    onChange={handleInput}
                                    id="name"
                                    className="border-2 w-[70%] ml-auto"
                                />
                            </div>
                            <div className="flex flex-row gap-2">
                                <label htmlFor="description" className="!text-black">Описание проекта</label>
                                <textarea 
                                    onChange={handleInput}
                                    className="min-h-50 w-[70%] border-2 ml-auto" 
                                    id="description"
                                >

                                </textarea>
                            </div>

                            <section className="flex flex-row w-full">
                                <div>
                                    <p>Дата начала проекта</p>
                                    <Calendar 
                                    onChange={setStartUpDate} 
                                    value={startUpDate}
                                    />
                                </div>
                                <div className="ml-auto">
                                    <p>Дата конца проекта</p>
                                    <Calendar 
                                    onChange={setEndDate}
                                    value={endDate}
                                    />
                                </div>
                            </section>

                            <button
                                className="min-h-10 w-50 flex bg-palette-color4 flex-col items-center justify-center self-center rounded-xl font-semibold text-xl text-white hover:cursor-pointer"
                            >
                                Создать проект
                            </button>
                        </form>
                    </div>
                </section>
                :
                <></>
            }

            <div className={`flex flex-row flex-1 mt-20 ${showPopUp ? "overflow-hidden" : ""}`}>
                <Sidebar show={true} choice={1} />
                <main className="w-[83%] ml-[17%] pr-5">
                    <h1 className="text-[#1D3266] text-4xl font-bold my-5">Мои проекты</h1>
                    <ul className="flex flex-row w-full gap-5 flex-wrap">
                        {projectsList?.map((item, index) => (
                            <li className="flex flex-col w-full">
                                <section key={index} 
                                    onClick={() => 
                                        {
                                            setProjectMaximize(prev => {
                                                const newMaximizedProjects = [...prev!];
                                                newMaximizedProjects[index] = !newMaximizedProjects[index];
                                                return newMaximizedProjects;
                                            });
                                        }
                                    } 
                                    className={`flex items-center p-5 w-full rounded-xl overflow-hidden ${maximizedProjects![index] ? "max-h-[40vh] flex-col" : "h-[5vh] flex-row"} border-3 border-palette-color5 mb-2`}
                                >
                                    <div className="">{item.name}</div>
                                    <div className={`${maximizedProjects![index] ? "" : "ml-auto mr-5"} flex flex-row`}>
                                        <div className="">{`${unixDateStampToDate(item.startTime).getDate()}/${unixDateStampToDate(item.startTime).getMonth() + 1}/${unixDateStampToDate(item.startTime).getFullYear()}→`}</div>
                                        <div className="">{`${unixDateStampToDate(item.endTime).getDate()}/${unixDateStampToDate(item.endTime).getMonth() + 1}/${unixDateStampToDate(item.endTime).getFullYear()}`}</div>
                                    </div>
                                    {
                                        maximizedProjects![index] ?
                                        <>
                                        <div className="overflow-auto">
                                            {item.description}
                                        </div>
                                        <button onClick={() => selectProject(item.id)} className={`${maximizedProjects![index] ? "mt-auto" : ""} bg-palette-color4 p-1 px-2 rounded-xl text-white z-50 hover:cursor-pointer`}>Перейти к проекту</button>
                                        </>
                                         :
                                        <>
                                        </>
                                    }
                                </section>
                                <button className="absolute right-8 flex justify-flex translate-y-[42%]" onClick={() => deleteProject(item.id)}>
                                    <span className="material-symbols-outlined select-none hover:text-red-700">
                                        close
                                    </span>
                                </button>
                            </li>
                            ))
                        }
                        <li key={0} className="flex flex-col w-full h-[5vh]">
                            <button
                                onClick={() => setShowPopUp(true)}
                                className="h-full w-full flex bg-palette-color4 flex-row items-center justify-center rounded-xl font-semibold text-xl text-white hover:cursor-pointer"
                            >
                                <div>Создать проект</div>
                                <span className="text-2xl ml-5">+</span>
                            </button>
                        </li>
                    </ul>
                </main>
            </div>
        </div>
    );
};

export default Projects;
