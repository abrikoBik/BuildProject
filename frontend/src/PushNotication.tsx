import { useState, useEffect } from "react";

interface PushNotificationProps {
    text : string
    duration? : number
}

const PushNotification = ({text, duration = 3000} : PushNotificationProps) => {

    const [isVisible, setVisibility] = useState(false);

    useEffect(() => {
        if(text)
        {
            setVisibility(true);

            const timer = setTimeout(() => {
                setVisibility(false);
            }, duration);

            return () => clearTimeout(timer);
        }
    }, [text, duration])

    if(!text) return null;

    return (
        <div className={`absolute ${isVisible ? "infoBoxAnimationIn" : "infoBoxAnimationOut"} w-[30%] h-[5%] m-auto left-0 right-0 z-50 rounded-b-xl border-3 border-t-0 flex justify-center items-center border-palette-color5 bg-palette-color2`}>
            <p id="infoBoxText">{text}</p>
        </div>
    )

}

export default PushNotification;