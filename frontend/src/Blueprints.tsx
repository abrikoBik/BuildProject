import Header from './Header';
import Sidebar from './Sidebar';
import Img1 from '../src/assets/image-1.png';
import Img2 from '../src/assets/image-2.png';
import Img3 from '../src/assets/image-3.png';
import Img4 from '../src/assets/image.png';

const Blueprints = () => {
    return (
        <div className="flex flex-col h-screen">
            <div className='mb-5'></div>
            <Header />
            <div className="flex flex-row flex-1">
                <Sidebar show={true} choice={3}/>
                <main className="w-[83%] ml-[17%] pr-5">
                    <h1 className="text-[#1D3266] text-4xl font-bold mt-5">Чертежи и планы</h1>
                    <div className="flex flex-row gap-10 mb-5">
                        <h3 className="text-lg">Заказчик: Иванов И. С.</h3>
                        <h3 className="text-lg">Объект: Московкая д. 68</h3>
                    </div>
                    <section className="w-full h-screen flex flex-row items-start gap-10">
                        <button
                            className="w-80 h-40 text-center border-2 bg-palette-color4 pt-1 rounded-xl font-semibold text-xl text-white hover:cursor-pointer"
                        >
                            <span className="material-symbols-outlined-huge">
                                upload
                            </span>
                            <div>Загрузить файлы</div>
                        </button>

                        <div className="flex flex-col w-100 border-3 rounded-xl p-5">
                            <img src={Img3} alt="img1" className='rounded-xl min-h-[50%]'/>
                            <section className="flex flex-row my-5 gap-5">
                                <div>
                                    <div className="text-xl font-bold mb-2">Чертеж фасада</div>
                                    <div>Фасад, общая высота 8 м</div>
                                </div>
                                <div className="flex flex-row gap-5 ml-auto justiyf-self-end">
                                    <button
                                        className="flex flex-row w-8 h-8 rounded-lg bg-palette-color3 items-center justify-center"
                                    >
                                        <span className="material-symbols-outlined">
                                            edit
                                        </span>
                                    </button>
                                    <button
                                        className="flex flex-row w-8 h-8 rounded-lg bg-palette-color3 items-center justify-center"
                                    >
                                        <span className="material-symbols-outlined">
                                            search
                                        </span>
                                    </button>
                                </div>
                            </section>
                        </div>

                        <div className="flex flex-col w-100 border-3 rounded-xl p-5">
                            <img src={Img4} alt="img1" className='rounded-xl min-h-[50%]'/>
                            <section className="flex flex-row my-5 gap-5">
                                <div>
                                    <div className="text-xl font-bold mb-2">План фундамента</div>
                                    <div>Фундамент, вид сверху</div>
                                </div>
                                <div className="flex flex-row gap-5 ml-auto justiyf-self-end">
                                    <button
                                        className="flex flex-row w-8 h-8 rounded-lg bg-palette-color3 items-center justify-center"
                                    >
                                        <span className="material-symbols-outlined">
                                            edit
                                        </span>
                                    </button>
                                    <button
                                        className="flex flex-row w-8 h-8 rounded-lg bg-palette-color3 items-center justify-center"
                                    >
                                        <span className="material-symbols-outlined">
                                            search
                                        </span>
                                    </button>
                                </div>
                            </section>
                        </div>

                        <div className="flex flex-col w-100 border-3 rounded-xl p-5">
                            <img src={Img2} alt="img1" className='rounded-xl min-h-[50%]'/>
                            <section className="flex flex-row my-5 gap-5">
                                <div>
                                    <div className="text-xl font-bold mb-2">Каркас дома</div>
                                    <div>Схема строительной системы</div>
                                </div>
                                <div className="flex flex-row gap-5 ml-auto justiyf-self-end">
                                    <button
                                        className="flex flex-row w-8 h-8 rounded-lg bg-palette-color3 items-center justify-center"
                                    >
                                        <span className="material-symbols-outlined">
                                            edit
                                        </span>
                                    </button>
                                    <button
                                        className="flex flex-row w-8 h-8 rounded-lg bg-palette-color3 items-center justify-center"
                                    >
                                        <span className="material-symbols-outlined">
                                            search
                                        </span>
                                    </button>
                                </div>
                            </section>
                        </div>

                        <div className="flex flex-col w-100 border-3 rounded-xl p-5">
                         <img src={Img1} alt="img1" className='rounded-xl min-h-[50%]'/>
                            <section className="flex flex-row my-5 gap-5">
                                <div>
                                    <div className="text-xl font-bold mb-2">План этажа</div>
                                    <div>Первый этаж, двери и мебель</div>
                                </div>
                                <div className="flex flex-row gap-5 ml-auto justiyf-self-end">
                                    <button
                                        className="flex flex-row w-8 h-8 rounded-lg bg-palette-color3 items-center justify-center"
                                    >
                                        <span className="material-symbols-outlined">
                                            edit
                                        </span>
                                    </button>
                                    <button
                                        className="flex flex-row w-8 h-8 rounded-lg bg-palette-color3 items-center justify-center"
                                    >
                                        <span className="material-symbols-outlined">
                                            search
                                        </span>
                                    </button>
                                </div>
                            </section>
                        </div>

                        

                        
                    </section>
                </main>
            </div>
        </div>
    );
}

export default Blueprints;
