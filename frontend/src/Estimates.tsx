import Header from './Header';
import Sidebar from './Sidebar';
import { useState } from 'react';

interface LineItem {
    name: string,
    unit: string,
    amount: number,
    price: number,
    cost: number
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

    return (
        <div className="flex flex-col h-screen">
            <Header />

            <div className="flex flex-row flex-1">
                <Sidebar show={true} choice={2}/>

                <main className="w-[83%] ml-[17%] pr-5">
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
                            <ul className="flex flex-row bold-child relative items-center z-0 rounded-t-xl mb-[-40px] w-full h-10 bg-palette-color2 px-[2%] justify-between">
                                <li>Наименование</li>
                                <li>Единица</li>
                                <li>Количество</li>
                                <li>Цена</li>
                                <li>Стоимость</li>
                            </ul>
                            <section className="flex flex-row w-full h-full rounded-xl border-3 py-[50px] z-50 border-palette-color4">
                                <form className="flex flex-col items-center w-full overflow-scroll">
                                    {lineArray.map((item, index) => {
                                        return (
                                            <section key={index} className="flex flex-row justify-between w-full px-[2%]">
                                                <div className="self-center">{item.name}</div>
                                                <div className="self-center">{item.unit}</div>
                                                <div className="self-center">{item.amount}</div>
                                                <div className="self-center">{item.price} р</div>
                                                <div className="self-center">{item.cost} р</div>
                                            </section>
                                        )
                                    })}
                                    <section className="flex flex-row justify-between w-full px-[2%]">
                                        <input type="text" onChange={handleChange} className="border-2 border-palette-color4 mb-2 text-black" id="name"/>
                                        {/* should be dropbox */}<input type="text" onChange={handleChange} className="border-2 border-palette-color4 mb-2" id="unit"/>
                                        <input onChange={handleChange} type="number" className="border-2 border-palette-color4 mb-2" id="amount"/>
                                        <input onChange={handleChange} type="number" className="border-2 border-palette-color4 mb-2" id="price"/>
                                        <div className="self-center" id="cost">{formData.cost} р</div>
                                    </section>
                                    <input onClick={handleAppending} type="reset" value="Добавить строку" className="w-40 h-10 hover:cursor-pointer bg-palette-color4 text-white rounded-lg"/>
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
                        <ul className="flex flex-col gap-5">
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
        </div>
    );
}

export default Estimates;