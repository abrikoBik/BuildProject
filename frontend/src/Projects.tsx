import Sidebar from './Sidebar';
import Header from './Header';
import { useState } from 'react';
import Calendar from 'react-calendar';
import 'react-calendar/dist/Calendar.css';
import { Value, UserRequest } from './types/projects'
import axios from 'axios';

const Projects = () => {
    const [userRequest, setUserRequest] = useState<UserRequest>({
        name: "",
        description: "",
        userID: 2,
        startTime: 1,
        endTime: 1
    });

    const [showPopUp, setShowPopUp] = useState(false);

    const [startUpDate, setStartUpDate] = useState<Value>(new Date());
    const [endDate, setEndDate] = useState<Value>(new Date());

    const handleInput = (e : any) => {
        setUserRequest({
            ...userRequest,
            [e.target.id]: e.target.value,
            userID: 2,
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

        } catch (err : any) {
            console.log(err);
        }
    }

    // Get list of projects

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
                        <li key={0} className="flex flex-col w-[20em] h-[30em]">
                            <button
                                onClick={() => setShowPopUp(true)}
                                className="h-full w-full flex bg-palette-color4 flex-col items-center justify-center rounded-xl font-semibold text-xl text-white hover:cursor-pointer"
                            >
                                <span className="text-9xl mb-2">+</span>
                                <div>Создать проект</div>
                            </button>
                        </li>

                    </ul>
                </main>
            </div>
        </div>
    );
};

export default Projects;
