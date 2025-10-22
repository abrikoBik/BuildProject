import Header from './Header';
import Sidebar from './Sidebar';
import { useEffect, useState } from 'react';
import PushNotification from './PushNotication';
import { LineItem, EstimateGroup } from './types/estimates';
import axios from 'axios';

/*const EditableCell = ({ value, onSave, onChange, textStyle, inputStyle }: EditableCellProps) => {
    const [isEdit, setEdit] = useState(false);
    const [curValue, setValue] = useState(value);

    const handleEditEnd = (event: React.KeyboardEvent<HTMLInputElement> | React.FocusEvent) => {
        if (event.type === 'keydown' && (event as React.KeyboardEvent).key !== 'Enter') {
            return;
        }

        setEdit(false);
        if (onSave) {
            onSave(curValue);
        }
    };

    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setValue(e.target.value);

        if(onChange) {
            onChange(e.target.value);
        }
    };

    return (
        <>
            {isEdit ? (
                <input
                    type="text"
                    value={curValue}
                    onChange={handleChange}
                    onBlur={handleEditEnd}
                    onKeyDown={handleEditEnd}
                    autoFocus
                    className={inputStyle}
                />
            ) : (
                <div 
                    onDoubleClick={() => setEdit(true)}
                    className={textStyle}
                >
                    {curValue}
                </div>
            )}
        </>
    );
};*/

const Estimates = () => {

    const [formData, setFormData] = useState<LineItem>({
        id: 0,
        name: '',
        unit: '',
        count: 0,
        cost: 0,
    });

    const [estimateAddWindow, setEstimateAddWindow] = useState(false);

    const [lineArray, setLineArray] = useState<LineItem[]>([]);

    const [pushNotificationText, setPushNotificationText] = useState("");

    const [estimateGroupsList, setEstimateGroupsList] = useState<EstimateGroup[]>();

    const [totalCost, setTotalCost] = useState(0);

    const [userRole, setUserRole] = useState<string>();

    const getUserRole = async () => {
        try {
            const roleRequest = await axios.get(
                `/api/users/${localStorage.getItem("userID")}/worker`
            );

            setUserRole(roleRequest.data.role);
        } catch (e : unknown) {
            setUserRole("client");
            console.log(e);
        }
    } 

    const pushNotification = (text : string) => {
        setPushNotificationText(text);
    }

    const handleAppending = async () => {
        if(!localStorage.getItem("selectedEstimateGroup")) {
            pushNotification("Группа смет не выбрана!");
            return;
        }
        try {
            const estimate = await axios.post(
                `/api/projects/${localStorage.getItem("selectedProjectID")}/estimateGroups/${localStorage.getItem("selectedEstimateGroup")}/estimates`,
                formData,
                {
                    headers: {
                        "Content-Type": "application/json"
                    }
                }
            )

            console.log(estimate);
            listEstimates();
            getTotalCost();
            pushNotification("Добавлена смета")
        } catch (e : unknown) {
            console.log(e);
        }
    };

    const handleInput = (event: React.ChangeEvent<HTMLInputElement>) => {
        setFormData({
            ...formData,
            [event.target.id]: event.target.id === "cost" || event.target.id === "count" ? parseInt(event.target.value) : event.target.value
        });
    }

    const listEstimates = async () => {
        try {
            const estimatesList = await axios.get(
                `api/projects/${localStorage.getItem("selectedProjectID")}/estimateGroups/${localStorage.getItem("selectedEstimateGroup")}/estimates`
            )

            console.log(estimatesList.data)
            setLineArray(estimatesList.data);
        } catch (e : unknown) {
            console.log(e)
        }
    }

    const listGroups = async () => {
        try {
            const estimateGroups = await axios.get(
                `api/projects/${localStorage.getItem("selectedProjectID")}/estimateGroups`
            )

            console.log(estimateGroups.data)
            setEstimateGroupsList(estimateGroups.data);
        } catch (e : unknown) {
            console.error(e);
        }
    }

    const deleteGroup = async (id : number) => {
        try {
            await axios.delete(
                `api/projects/${localStorage.getItem("selectedProjectID")}/estimateGroups/${id}`
            );
            listGroups();
            listEstimates();
            setLineArray(null!);
            localStorage.removeItem("selectedEstimateGroup");
            localStorage.removeItem("estimateGroup");
        } catch (e : unknown) {
            console.error(e);
        }
    }

    const addGroup = async (name : string) => {
        try {
            const estimateGroup = await axios.post(
                `api/projects/${localStorage.getItem("selectedProjectID")}/estimateGroups`,
                {"name": name},
                {
                    headers: {
                        "Content-Type": "application/json"
                    }
                }          
            )

            console.log(estimateGroup.data);
            listGroups();
            getTotalCost();
            pushNotification("Добавлена группа смет")
        } catch (e : unknown) {
            console.error(e);
        }
    }

    const getTotalCost = async () => {
        try {
            const estimateGroup = await axios.get(
                `api/projects/${localStorage.getItem("selectedProjectID")}/estimateGroups/${localStorage.getItem("selectedEstimateGroup")}`
            )

            setTotalCost(estimateGroup.data.totalCost);
        } catch (e : unknown) {
            console.error(e);
        }
    }

    const handleEstimateDelete = async (id: number) => {
        try {
            const estimate = await axios.delete(
                `api/projects/${localStorage.getItem("selectedProjectID")}/estimateGroups/${localStorage.getItem("selectedEstimateGroup")}/estimates/${id}`
            )

            console.log(estimate.data);
            listEstimates();
            getTotalCost();
            pushNotification("Удалена смета");
        } catch (e : unknown) {
            console.error(e);
        }
    }

    useEffect(() => {
        listEstimates();
        listGroups();
        getTotalCost();
        getUserRole();
    }, [])

    return (
        <main className="flex flex-col h-screen">
            <Header />
            <PushNotification text={pushNotificationText}/>
            <div className="flex flex-row flex-1">
                <Sidebar show={true} choice={2}/>
                { localStorage.getItem("selectedProjectID") ?
                <section className="w-[83%] ml-[17%] mt-20 pr-5">
                    {/* Selected estimate group */}
                    <section className="w-full flex-col justify-between h-1/2">
                        <section className="w-full h-20 flex flex-row">
                            <div>
                                <h1 className="text-[#1D3266] font-bold text-4xl">Сметы</h1>
                                <div className="text-md font-light">{ `Текущая группа смет: ${localStorage.getItem("estimateGroup") ? localStorage.getItem("estimateGroup") : "не выбрана"}` }</div>
                            </div>
                            { userRole === "client" ?
                                <></>
                                :
                                <>
                                {!estimateAddWindow ?
                                    <input onClick={() => setEstimateAddWindow(!estimateAddWindow)} type="button" value="Создать группу смет" className="w-60 h-15 hover:cursor-pointer text-center bg-palette-color4 pt-1 rounded-xl font-semibold text-xl text-white self-center ml-auto" />
                                    :
                                    <section className="mr-5 ml-auto my-auto">
                                        <input 
                                            type="text"
                                            id="estimateGroupName"
                                            autoFocus
                                            className="!border-3 mr-5"
                                            onKeyDown={(e) => {
                                                if(e.key == "Enter") {
                                                    setEstimateAddWindow(!estimateAddWindow);
                                                    addGroup(e.currentTarget.value);
                                                }
                                            }}
                                        />
    
                                        <label htmlFor="estimateGroupName" className="!text-black">
                                            <button 
                                                onClick={() => {
                                                    setEstimateAddWindow(!estimateAddWindow);
                                                    addGroup((document.getElementById("estimateGroupName") as HTMLInputElement).value);
                                                }} 
                                                className="bg-palette-color4 text-white rounded-lg p-1 px-2">
                                                    Добавить группу
                                            </button>
                                        </label>
                                    </section>
                                }
                                </>
                            }
                        </section>

                        <section className="flex flex-col h-[100%] w-full p-[3px] min-h-100">
                             <ul className="grid grid-cols-5 bold-child relative z-0 rounded-t-xl w-full h-10 bg-palette-color2 px-[2%] items-center mb-[-40px]">
                              <li className="text-center">Наименование</li>
                              <li className="text-center">Единица</li>
                              <li className="text-center">Количество</li>
                              <li className="text-center">Цена</li>
                              <li className="text-center">Стоимость</li>
                            </ul>
                            <section className="flex flex-row w-full h-full rounded-xl border-3 py-[50px] z-50 border-palette-color4">
                                <ul className="flex flex-col w-full h-full overflow-scroll">
                                    {
                                        lineArray != null ? lineArray.map((item, index) => (
                                            <li key={index} className="flex flex-row">
                                                <section className="flex flex-row w-full px-[2%] mb-5 justify-between">
                                                    {/* TODO: fix sizes */}
                                                    <div className="text-center w-[16%] h-[40px] border-2 border-white">{item.name ? item.name : "-"}</div>
                                                    <div className="text-center w-[16%] h-[40px] border-2 border-white">{item.unit ? item.unit : "-"}</div>
                                                    <div className="text-center w-[16%] h-[40px] border-2 border-white">{item.count}</div>
                                                    <div className="text-center w-[16%] h-[40px] border-2 border-white">{item.cost}</div>
                                                    <div className="text-center w-[16%]">{item.count * item.cost} р</div>
                                                </section>
                                                { userRole === "client" ?
                                                    <></>
                                                    :
                                                    <span onClick={() => handleEstimateDelete(item.id)} className="asbolute right-4 material-symbols-outlined select-none hover:text-red-700">
                                                        close
                                                    </span>
                                                }
                                            </li>
                                        ))
                                        :
                                        <>
                                        </>
                                    }
                                    { userRole === "client" 
                                        ?   
                                        <></>
                                        :
                                        <form className="flex flex-col items-center w-full">
                                            <section className="flex flex-row justify-between w-full px-[2%] mb-5">
                                                <input id="name" type="text" onChange={handleInput} className="!w-[15%] border-2 border-palette-color4 mb-2 text-black"/>
                                                <input id="unit" type="text" onChange={handleInput} className="!w-[15%] border-2 border-palette-color4 mb-2"/>
                                                <input id="count" onChange={handleInput} type="text" className="!w-[15%] border-2 border-palette-color4 mb-2"/>
                                                <input id="cost" onChange={handleInput} type="text" className="!w-[15%] border-2 border-palette-color4 mb-2"/>
                                                <div className="self-center w-[15%] text-center" id="cost">{formData.count * formData.cost} р</div>
                                            </section>
                                            <input onClick={handleAppending} type="button" value="Добавить строку" className="w-40 h-10 min-h-10 hover:cursor-pointer bg-palette-color4 text-white rounded-lg"/>
                                        </form>
                                    }
                                </ul>
                            </section>
                            <ul className="flex flex-row bold-child relative items-center z-0 rounded-b-xl mt-[-40px] w-full h-10 border-2 border-x-0 border-b-0 border-palette-color4 px-[2%]">
                                <li className="text-[#1D3266]">Итого</li>
                                <li className="ml-auto justify-self-end">{totalCost} р</li>
                            </ul>
                        </section>
                    </section>

                    {/* Estimate groups */}
                    <section className="w-full h-1/2 mb-5 mt-40">
                        <h2 className="text-[#1D3266] font-bold text-2xl mb-6">Группы смет</h2>
                        <ul className="flex flex-col gap-5 select-none">
                            { estimateGroupsList?.map((item, index) => (
                                <li key={index}>
                                    <div 
                                    onClick={() => {
                                        localStorage.setItem("selectedEstimateGroup", item.id.toString());
                                        localStorage.setItem("estimateGroup", item.name as string);
                                        getTotalCost();
                                        listEstimates();
                                    }}
                                    className={`flex flex-row h-18 border-3 justify-center rounded-xl p-3 hover:shadow-[0px_0px_5px_#37515fff] border-palette-color4 ${parseInt(localStorage.getItem("selectedEstimateGroup")!) === item.id ? "bg-[#C5C9D3]" : ""}`}
                                    >
                                        <h5 className="self-center">{item.name}</h5>                                    
                                        <span onClick={() => deleteGroup(item.id)} className="material-symbols-outlined absolute right-10 translate-y-[35%] hover:text-red-700">
                                            close
                                        </span>
                                    </div>
                                </li>
                            ))
                            }
                        </ul>
                    </section>
                </section>
                :
                <div className="flex w-full h-full justify-center items-center">Проект не выбран!</div>
                }
            </div>
        </main>
    );
}

export default Estimates;