export const authURL = process.env.NODE_ENV == "production" ? "/auth" : "http://localhost:8082";
export const searchURL = process.env.NODE_ENV == "production" ? "/search" : "http://localhost:8081";
export const dataURL = process.env.NODE_ENV == "production" ? "/data" : "http://localhost:8083";