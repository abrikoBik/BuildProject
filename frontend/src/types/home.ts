export interface TaskGroup {
    id: number;
    name: string | null;
}

export interface TaskRequest {
    startTime: number;
    endTime: number;
    name: string;
}

export interface TaskGet {
    id: number;
    name: string;
    startTime: number;
    endTime: number;
    status: boolean;
}