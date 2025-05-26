import Header from './Header';
import Sidebar from './Sidebar'

const Estimates = () => {
    return (
        <div className="flex flex-col h-screen">
            <Header />
            
            <div className="flex flex-row flex-1">
                <Sidebar show={true} choice={2}/>

                <main className="w-[83%] ml-[17%] pr-5">
                    <section className="w-full h-1/2">
                        <section className="w-full h-20 flex flex-row">
                            <div>
                                <h1 className="text-[#1D3266] font-bold text-4xl">Сметы</h1>
                                <div className="text-md font-light"> Заказчик: Иванов И. С.</div>
                            </div>
                            <input type="button" value="Создать смету" className="w-60 h-15 text-center border-2 bg-palette-color4 pt-1 rounded-xl font-semibold text-xl text-white self-center ml-auto justiyf-self-end" />
                        </section>

                        <section className="flex flex-col rounded-xl border-3 border-palette-color4 h-[70%]">
                            <div className="flex flex-row">

                            </div>
                        </section>
                    </section>

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
                        </ul>
                    </section>
                </main>
            </div>
        </div>
    );
}

export default Estimates;