//TODO : refactor to type folders
export type ToastStatusType = "success" | "error" | "warning" | "info";

export const ToastStatus = {
    SUCCESS: "success" as ToastStatusType,
    ERROR: "error" as ToastStatusType,
    WARNING: "warning" as ToastStatusType,
    INFO: "info" as ToastStatusType,
}