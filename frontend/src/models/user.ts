export type LoginDTO =  {
    username: string;
    password: string;
}

export type RegisterDTO = {
    username: string;
    password: string;
    email: string;
    role: string;
}

export type ChangePasswordDTO = {
    username: string;
    oldPassword: string;
    newPassword: string;
}