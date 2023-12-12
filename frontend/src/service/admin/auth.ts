import axios from "axios";

export const login = async (username: string, password: string) => {
    try {
        // const path = import.meta.env.VITE_AUTH_API_URL; //* For production
        const path = "http://localhost:8080"; //* For local development
        const response = await axios.post(`${path}/login`, {
            username,
            password,
        });
        return response.data;
    } catch (error) {
        console.error("Error:", error);
        return null;
    }
}