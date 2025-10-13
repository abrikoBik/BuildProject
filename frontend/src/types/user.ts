export interface UserRequest {
    login: string;
    name: string;
    phoneNumber: number;
    email: string;
    password: string;
}

export interface UserResponse {
    email: string;
    phoneNumber: number;
    login: string;
    name: string;
    registeredAt: string;
    id: number;
}
