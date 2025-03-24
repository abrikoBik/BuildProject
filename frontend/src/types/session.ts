export interface SessionRequest {
    login: string;
    password: string;
}

export interface SessionResponse {
    token: string;
    userID: number;
}
