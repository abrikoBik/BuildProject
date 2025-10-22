import Sidebar from './Sidebar';
import Header from './Header';
import { useState, useEffect, ChangeEvent } from 'react';
import { ProjectRequest } from './types/projects'

import 'react-calendar/dist/Calendar.css';
import axios from 'axios';
import { Map, Marker } from 'pigeon-maps'
import { osm } from 'pigeon-maps/providers'

const Projects = () => {
    let userID = 0;

    const [userRole, setUserRole] = useState({
        id: -1,
        role: ""
    });
    const [clientID, setClientID] = useState<number>();

    const getUserRole = async () => {
        try {
            const roleRequest = await axios.get(
                `/api/users/${localStorage.getItem("userID")}/worker`
            );

            setUserRole(roleRequest.data);
        } catch (e : unknown) {
            setUserRole({id:-1, role:"client"});
            console.log(e);
        }
    } 

    try {
        userID = JSON.parse(localStorage.getItem("user") as string).userID;
    } catch (e : unknown) {
        console.log(e);
        window.location.href = "/login";
    }

    const [userRequest, setUserRequest] = useState<ProjectRequest>({
        name: "",
        description: "",
        userID: userID,
        startTime: 1,
        endTime: 1,
        latitude: 0,
        longitude: 0,
        id: 0
    });

    const [showPopUp, setShowPopUp] = useState(false);

    const [Longitude, setLongitude] = useState<number>(0);
    const [Latitude, setLatitude] = useState<number>(0);
    const [AddressString, setAddress] = useState<string>();
    
    const [projectsList, setProjectsList] = useState<ProjectRequest[]>();

    const [maximizedProjects, setProjectMaximize] = useState<boolean[]>();

    const handleInput = (e : React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement | HTMLDataElement>) => {
        if(e.target.id != "startDate" && e.target.id != "endDate")
            setUserRequest({
                ...userRequest,
                [e.target.id]: e.target.value,
                userID: parseInt(localStorage.getItem("userID")!),
            })
        else if (e.target.id === "startDate")
            setUserRequest({
                ...userRequest,
                startTime: Math.floor((new Date(e.target.value)).getTime() / 1000),
            })
        else if (e.target.id === "endDate")
            setUserRequest({
                ...userRequest,
                endTime: Math.floor((new Date(e.target.value)).getTime() / 1000),
            })
    }

    const createProject = async () => {
        try {
            
            console.log(userRequest);
            console.log("Client ID", clientID);

            await axios.post(
                '/api/projects',
                {
                    ...userRequest,
                    userID: clientID
                },
                {
                    headers: {
                        "Content-Type": "application/json"
                    }
                }
            ).then(async (p) => {
                    console.log(p.data)
                    await axios.put(
                        `/api/projects/${p.data.id}/workers/${userRole.id}`,
                    )
                }
            )

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

            localStorage.removeItem("selectedProjectID");

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
            setProjectMaximize(Array.from({length: data?.length}, () => false))

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
        localStorage.removeItem("estimateGroup");
        localStorage.removeItem("selectedEstimateGroup");

        window.location.href = '/';
    }

    const handleCoordinates = async () => {
        try {
            console.log("Fetching coordinates...");
            const response = await axios.get(
                `https://nominatim.openstreetmap.org/search?q=${encodeURIComponent(AddressString!)}&format=json`
            );

            const results = response.data;
            if (!Array.isArray(results) || results.length === 0) {
                console.warn("No coordinates found for address:", AddressString);
                return;
            }

            const firstResult = results[0];
            const lat = parseFloat(firstResult.lat);
            const lon = parseFloat(firstResult.lon);

            setLatitude(lat);
            setLongitude(lon);
            
            setUserRequest({...userRequest, latitude: lat, longitude: lon});

            console.log(`Address is ${AddressString}, lat: ${lat}, lon: ${lon}`);

        } catch (e) {
            if (e instanceof Error) {
                console.error("Error fetching coordinates:", e.message);
            } else {
                console.error("Unknown error:", e);
            }
        }
    };

    useEffect(() => {
        listProjects();
        getUserRole();
    }, [])

    return (
        <div className="flex flex-col h-screen">
            <Header />

            { showPopUp ? 
                <section className="fixed left-0 top-0 w-screen h-screen bg-[rgba(0,0,0,0.4)] z-150">
                    <div className="flex flex-col relative left-1/2 top-1/2 -translate-1/2 w-[60%] h-[80%] rounded-xl border-2 border-palette-color4 bg-white overflow-y-scroll">
                        <button onClick={() => setShowPopUp(false)} type="button" className="hover:cursor-pointer hover:text-red-700 w-7 h-7 flex items-center justify-center self-end">
                            <span className="material-symbols-outlined">
                                close
                            </span>
                        </button>

                        <button
                            onClick={createProject}
                            className="min-h-10 w-50 flex bg-palette-color4 flex-col items-center justify-center self-center rounded-xl font-semibold text-xl text-white hover:cursor-pointer"
                        >
                                Создать проект
                        </button>

                        <section className="flex flex-col w-full h-full p-5 gap-5">
                            <div className="flex flex-row gap-2">
                                <label htmlFor="id" className="!text-black">ID пользователя</label>
                                <input 
                                    onChange={(e) => setClientID(parseInt(e.currentTarget.value))}
                                    id="id"
                                    className="border-2 w-[70%] ml-auto"
                                />
                            </div>
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
                                    <input id="startDate" type="date" onChange={handleInput}/>
                                </div>
                                <div className="ml-auto">
                                    <p>Дата конца проекта</p>
                                    <input id="endDate" type="date" onChange={handleInput}/>
                                </div>
                            </section>

                            <div>
                                <div>Карта</div>
                                <Map
                                    provider={osm}
                                    height={200}
                                    defaultCenter={[47.2153657, 38.9285216]}
                                    center={[Latitude, Longitude]}
                                    defaultZoom={11}
                                >
                                    <Marker
                                        width={50}
                                        anchor={[Latitude, Longitude]} 
                                    />
                                </Map>
                                <section className="w-full">
                                    <input 
                                        type="text"
                                        onChange={(e : ChangeEvent<HTMLInputElement>) => setAddress(e.target.value)}
                                        onKeyDown={(e) => {
                                            if(e.key === "Enter")
                                            {
                                                setAddress(e.currentTarget.value);
                                                handleCoordinates();
                                            }
                                        }}
                                        className="w-[80%]"
                                    />
                                    <button 
                                        onClick={handleCoordinates}
                                        type="button"
                                        className="bg-palette-color4 text-white rounded-xl p-2 w-[20%]"
                                    >
                                    Добавить метку
                                    </button>
                                </section>
                            </div>
                        </section>
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
                            <li key={index} className="flex flex-col w-full">
                                <section
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
                                        <Map
                                            provider={osm}
                                            height={200}
                                            defaultCenter={[47.2153657, 38.9285216]}
                                            center={[item.latitude, item.longitude]} 
                                            defaultZoom={15}
                                        >
                                            <Marker
                                                width={50}
                                                anchor={[item.latitude, item.longitude]} 
                                            />
                                        </Map>
                                        <button onClick={() => selectProject(item.id)} className={`${maximizedProjects![index] ? "mt-auto" : ""} bg-palette-color4 p-1 px-2 rounded-xl text-white z-50 hover:cursor-pointer`}>Перейти к проекту</button>
                                        </>
                                         :
                                        <>
                                        </>
                                    }
                                </section>
                                { userRole.role === "client" ?
                                    <></>
                                    :
                                    <button className="absolute right-8 flex justify-flex translate-y-[42%]" onClick={() => deleteProject(item.id)}>
                                        <span className="material-symbols-outlined select-none hover:text-red-700">
                                            close
                                        </span>
                                    </button>
                                }
                            </li>
                            ))
                        }
                        { userRole.role === "client" ?
                            <></>
                            :
                            <li key={0} className="flex flex-col w-full h-[5vh]">
                                <button
                                    onClick={() => setShowPopUp(true)}
                                    className="h-full w-full flex bg-palette-color4 flex-row items-center justify-center rounded-xl font-semibold text-xl text-white hover:cursor-pointer min-h-10"
                                >
                                    <div>Создать проект</div>
                                    <span className="text-2xl ml-5">+</span>
                                </button>
                            </li>
                        }
                    </ul>
                </main>
            </div>
        </div>
    );
};

export default Projects;
