export interface UserRequest {
    login: string;
    name: string;
    email: string;
    password: string;
}

export interface UserResponse {
    email: string;
    login: string;
    name: string;
    registeredAt: string;
}
