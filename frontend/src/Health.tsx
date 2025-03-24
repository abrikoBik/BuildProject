import { useState, useEffect } from "react";
import axios from "axios";

export const Health = () => {
    const [checkHealth, setCheck] = useState(false);

    useEffect(() => {
        const apiURL = "http://0.0.0.0:3000/health";
        axios.get(apiURL).then((resp) => {
            setCheck(resp.status == 200);
        });
    }, [setCheck]);

    return (
        <>
            <section className="absolute top-[50%] left-[50%] translate-x-[-50%] translate-y-[-50%]">
                <div className="text-8xl font-bold">
                    {checkHealth ? "✅ OKAY" : "❌ SERVER DON'T RESPONSE"}
                </div>
            </section>
        </>
    );
};
