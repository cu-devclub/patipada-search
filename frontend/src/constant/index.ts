export type ToastStatusType = "success" | "error" | "warning" | "info";

export const ToastStatus = {
    SUCCESS: "success" as ToastStatusType,
    ERROR: "error" as ToastStatusType,
    WARNING: "warning" as ToastStatusType,
    INFO: "info" as ToastStatusType,
}

export const PASSWORD_REQUIRED_LENGTH=8
 

export const Role = {
    USER: "user",
    ADMIN: "admin",
    SUPER_ADMIN: "super_admin",
}

export const SERVER_ERROR_MESSAGE = {
  // message match with server response
  USERNAME_ALREADY_EXISTS: "Username already exists",
  EMAIL_ALREADY_EXISTS: "Email already exists",
};

export const ERR_Messages = {
  SYSTEM_ERROR: "ระบบมีปัญหา โปรดลองใหม่อีกครั้งในภายหลัง",
  INVALID_USERNAME_OR_PASSWORD: "ชื่อผู้ใช้หรือรหัสผ่านไม่ถูกต้อง",
  FILL_ALL_FIELDS: "กรุณากรอกข้อมูลให้ครบถ้วน",
  NOT_ATTACH_TOKEN: "ไม่พบ session token",
  NO_PERMISSION_REGISTER: "ไม่มีสิทธิ์ในการสมัครสมาชิกในตำแหน่งดังกล่าว",
  INVALID_EMAIL: "รูปแบบอีเมลไม่ถูกต้อง",
  NOT_FOUND_EMAIL: "ไม่พบอีเมลนี้ในระบบ",
  NOT_FOUND_TOKEN: "ไม่พบ token นี้ในระบบ",
  INVALID_PASSWORD_FORMAT: `รหัสผ่านต้องมีความยาวอย่างน้อย ${PASSWORD_REQUIRED_LENGTH} ตัวอักษร`,
  SAME_PASSWORD: "รหัสผ่านเดิมและรหัสผ่านใหม่ตรงกัน",
  INVALID_OLD_PASSWORD: "รหัสผ่านเดิมไม่ถูกต้อง",
};

// Mapping user-friendly error messages to error message from server
export const ERR_Messages_MAP: { [key: string]: string } = {
  [SERVER_ERROR_MESSAGE.USERNAME_ALREADY_EXISTS]:
    "มีผู้ใช้งานชื่อผู้ใช้นี้แล้ว",
  [SERVER_ERROR_MESSAGE.EMAIL_ALREADY_EXISTS]: "มีผู้ใช้งานอีเมลนี้แล้ว",
};
