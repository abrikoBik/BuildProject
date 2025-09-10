import Sidebar from './Sidebar';
import Header from './Header';
import { useState } from 'react';
import Calendar from 'react-calendar';
import 'react-calendar/dist/Calendar.css';

type ValuePiece = Date | null;

type Value = ValuePiece | [ValuePiece, ValuePiece];

const Projects = () => {
    const projects = Array.from({ length: 50 }, (_, i) => i + 1);

    const [showPopUp, setShowPopUp] = useState(false);

    const [startUpDate, setStartUpDate] = useState<Value>(new Date());
    const [endDate, setEndDate] = useState<Value>(new Date());

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

                        <form action="submit" className="flex flex-col w-full h-full p-5 gap-5">
                            <div className="flex flex-row gap-2">
                                <label htmlFor="name" className="!text-black">Название проекта</label>
                                <input id="name" className="border-2 w-[70%] ml-auto" />
                            </div>
                            <div className="flex flex-row gap-2">
                                <label htmlFor="description" className="!text-black">Описание проекта</label>
                                <textarea className="min-h-50 w-[70%] border-2 ml-auto" id="description"></textarea>
                            </div>

                            <section className="flex flex-row w-full">
                                <div>
                                    <p>Дата начала проекта</p>
                                    <Calendar onChange={setStartUpDate} value={startUpDate}/>
                                </div>
                                <div className="ml-auto">
                                    <p>Дата конца проекта</p>
                                    <Calendar onChange={setEndDate} value={endDate}/>
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

                        {/* Projects dipslay */}
                        {projects.map((i) => (
                            <li key={i + 1} className="flex flex-col w-[20em] h-[30em]">
                                <div
                                    style={{
                                        backgroundColor: `rgba(${Math.min(200 + i, 255)}, 0, ${Math.min(50 + i, 255)}, 1)`
                                    }}
                                    className="w-full h-[70%] border-3 rounded-xl flex items-center justify-center"
                                >
                                    image
                                </div>
                                <h3 className="w-full text-center text-xl my-2">{`Проект ${i}`}</h3>
                                <button
                                    className="w-full h-15 text-center border-2 bg-palette-color4 pt-1 rounded-xl font-semibold text-xl text-white self-center"
                                >
                                    Перейти к проекту
                                </button>
                            </li>
                        ))}
                    </ul>
                </main>
            </div>
        </div>
    );
};

export default Projects;
