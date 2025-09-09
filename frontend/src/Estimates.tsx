import Header from './Header';
import Sidebar from './Sidebar';
import { useState } from 'react';
import PushNotification from './PushNotication';

interface LineItem {
    name: string,
    unit: string,
    amount: number,
    price: number,
    cost: number
};

interface EditableCellProps {
    value: string;
    onSave?: (newValue: string) => void;
    onChange?: (newValue: string) => void;
}

const EditableCell = ({ value, onSave, onChange }: EditableCellProps) => {
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
        <div className="text-center">
            {isEdit ? (
                <input
                    type="text"
                    value={curValue}
                    onChange={handleChange}
                    onBlur={handleEditEnd}
                    onKeyDown={handleEditEnd}
                    autoFocus
                    className="w-full px-2 border-2"
                />
            ) : (
                <div 
                    onDoubleClick={() => setEdit(true)}
                    className="cursor-text hover:bg-gray-100 p-1"
                >
                    {curValue}
                </div>
            )}
        </div>
    );
};

const Estimates = () => {
    const [formData, setFormData] = useState<LineItem>({
      name: '',
      unit: '',
      amount: 0,
      price: 0,
      cost: 0
    });

    const [lineArray, setLineArray] = useState<LineItem[]>([]);
    const [totalCost, setTotalCost] = useState(0);

    const [pushNotificationText, setPushNotificationText] = useState("");

    const pushNotification = (text : string) => {
        setPushNotificationText(text);
    }

    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
      const { id, value } = e.target;
      const numericValue = id === 'name' || id === 'unit' ? value : parseFloat(value) || 0;
    
      setFormData(prev => {
        const newData = { ...prev, [id]: numericValue };
        if (id === 'price' || id === 'amount') {
            newData.cost = newData.price * newData.amount;
        }
        return newData;
      });
    };

    const handleAppending = () => {
        setLineArray(prev => [...prev, formData]);
        setTotalCost(totalCost + formData.cost);
        setFormData({
          name: '',
          unit: '',
          amount: 0,
          price: 0,
          cost: 0
        });
    };

    const handleDelete = (index : number) => {
        if(index != null)
        {
            const array = [...lineArray];
            setTotalCost(totalCost - array[index].cost);
            array.splice(index, 1);
            setLineArray(array);
            pushNotification(`Удалена строка номер ${index+1}`);
        }
    }

    const recalculateTotalCost = () => {
        setTotalCost(0);
        lineArray.map((item) => {
            setTotalCost(totalCost + item.cost);
        })  
    }

    {/* TODO: fix buttons size when add new estimation line */}
    return (
        <main className="flex flex-col h-screen">
            <Header />
            <PushNotification text={pushNotificationText}/>
            <div className="flex flex-row flex-1">
                <Sidebar show={true} choice={2}/>

                <main className="w-[83%] ml-[17%] mt-20 pr-5">
                    {/* Selected estimate group */}
                    <section className="w-full h-[60%]">
                        <section className="w-full h-20 flex flex-row">
                            <div>
                                <h1 className="text-[#1D3266] font-bold text-4xl">Сметы</h1>
                                <div className="text-md font-light"> Заказчик: Иванов И. С.</div>
                            </div>
                            <input type="button" value="Создать смету" className="w-60 h-15 hover:cursor-pointer text-center bg-palette-color4 pt-1 rounded-xl font-semibold text-xl text-white self-center ml-auto justiyf-self-end" />
                        </section>

                        <section className="flex flex-col h-[80%] w-full p-[3px]">
                             <ul className="grid grid-cols-5 bold-child relative z-0 rounded-t-xl w-full h-10 bg-palette-color2 px-[2%] items-center mb-[-40px]">
                              <li className="text-center">Наименование</li>
                              <li className="text-center">Единица</li>
                              <li className="text-center">Количество</li>
                              <li className="text-center">Цена</li>
                              <li className="text-center">Стоимость</li>
                            </ul>
                            <section className="flex flex-row w-full h-full rounded-xl border-3 py-[50px] z-50 border-palette-color4">
                                <form className="flex flex-col items-center w-full overflow-scroll">
                                    {lineArray.map((item, index) => {
                                        return (
                                        <div 
                                            key={index}
                                            className="grid grid-cols-5 gap-4 w-full px-[2%] py-3 items-center border-b border-gray-200"
                                        >
                                            {/* TODO: fix bug when input stays blank */}
                                            <EditableCell 
                                              value={item.name || '-'} 
                                              onSave={(newValue) => {
                                                const updated = [...lineArray];
                                                updated[index].name = newValue ?? '-';
                                                setLineArray(updated);
                                              }}
                                            />
                                            <EditableCell 
                                              value={item.unit || '-'} 
                                              onSave={(newValue) => {
                                                const updated = [...lineArray];
                                                updated[index].unit = newValue ?? '-';
                                                setLineArray(updated);
                                              }}
                                            />
                                            <EditableCell 
                                              value={String(item.amount ?? '0')} 
                                              onSave={(newValue) => {
                                                const updated = [...lineArray];
                                                updated[index].amount = parseFloat(newValue) ?? 0;
                                                recalculateTotalCost();
                                                setLineArray(updated);
                                              }}
                                              onChange={(newValue) => {
                                                const updated = [...lineArray];
                                                updated[index].cost = (parseFloat(newValue) ?? 0) * updated[index].price;
                                                setLineArray(updated);
                                              }}
                                            />
                                            <EditableCell 
                                              value={String(item.price ?? '0')} 
                                              onSave={(newValue) => {
                                                const updated = [...lineArray];
                                                updated[index].price = parseFloat(newValue) ?? 0;
                                                recalculateTotalCost();
                                                setLineArray(updated);
                                              }}
                                              onChange={(newValue) => {
                                                const updated = [...lineArray];
                                                updated[index].cost = (parseFloat(newValue) ?? 0) * updated[index].amount;
                                                setLineArray(updated);
                                              }}
                                            />
                                            <div className="flex flex-row gap-10 justify-center pl-15">
                                                <div className="text-center font-medium">{item.cost ?? '0'} р</div>
                                                <div onMouseDown={() => handleDelete(index)} 
                                                    className="flex w-7 h-7 items-center justify-center hover:text-red-700 select-none hover:cursor-pointer"
                                                    aria-label="Удалить строку"
                                                >
                                                    <span className="material-symbols-outlined">
                                                        close
                                                    </span>
                                                </div>
                                            </div>
                                        </div>
                                        )
                                    })}

                                    <section className="flex flex-row justify-between w-full px-[2%]">
                                        <input type="text" onChange={handleChange} className="border-2 border-palette-color4 mb-2 text-black" id="name"/>
                                        {/* should be dropbox */}<input type="text" onChange={handleChange} className="border-2 border-palette-color4 mb-2" id="unit"/>
                                        <input onChange={handleChange} type="text" className="border-2 border-palette-color4 mb-2" id="amount"/>
                                        <input onChange={handleChange} type="text" className="border-2 border-palette-color4 mb-2" id="price"/>
                                        <div className="self-center" id="cost">{formData.cost} р</div>
                                    </section>
                                    <input onClick={handleAppending} type="reset" value="Добавить строку" className="w-40 h-10 min-h-10 hover:cursor-pointer bg-palette-color4 text-white rounded-lg"/>
                                </form>
                            </section>
                            <ul className="flex flex-row bold-child relative items-center z-0 rounded-b-xl mt-[-40px] w-full h-10 border-2 border-x-0 border-b-0 border-palette-color4 px-[2%]">
                                <li className="text-[#1D3266]">Итого</li>
                                <li className="ml-auto justify-self-end">{totalCost} р</li>
                            </ul>
                        </section>
                    </section>

                    {/* Estimate groups */}
                    <section className="w-full h-1/2 mb-5">
                        <h2 className="text-[#1D3266] font-bold text-2xl mb-6">Предварительное ведение смет</h2>
                        <ul className="flex flex-col gap-5 select-none">
                            <li>
                                <div className="flex flex-row h-18 border-3 rounded-xl p-3 hover:shadow-[0px_0px_5px_#37515fff] border-palette-color4 ">
                                    <h5 className="self-center">Фундамент</h5>                                    
                                    <div className="ml-auto self-center">25 апр - 2025</div>
                                </div>
                            </li>
                            <li>
                                <div className="flex flex-row h-18 border-3 rounded-xl p-3 hover:shadow-[0px_0px_5px_#37515fff] border-palette-color4 ">
                                    <h5 className="self-center">Пол</h5>                                    
                                    <div className="ml-auto self-center">19 апр - 2025</div>
                                </div>
                            </li>
                            <li>
                                <div className="flex flex-row h-18 border-3 rounded-xl p-3 hover:shadow-[0px_0px_5px_#37515fff] border-palette-color4 ">
                                    <h5 className="self-center">Стены</h5>                                    
                                    <div className="ml-auto self-center">15 апр - 2025</div>
                                </div>
                            </li>
                            <li>
                                <div className="flex flex-row h-18 border-3 rounded-xl p-3 hover:shadow-[0px_0px_5px_#37515fff] border-palette-color4 ">
                                    <h5 className="self-center">Терраса</h5>                                    
                                    <div className="ml-auto self-center">10 апр - 2025</div>
                                </div>
                            </li>
                            <li className="w-full h-5">

                            </li>
                        </ul>
                    </section>
                </main>
            </div>
        </main>
    );
}

export default Estimates;