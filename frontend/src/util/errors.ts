import axios from "axios";

export default {
    axiosError: (err: unknown) => {
        if (axios.isAxiosError(err)) {
            alert(err?.response?.data);
        }
        console.error(err);
    },
};
