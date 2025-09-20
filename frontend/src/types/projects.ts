export type ValuePiece = Date | null;

export type Value = ValuePiece | [ValuePiece, ValuePiece];

export interface UserRequest  {
    name: string,
    userID: number,
    description: string,
    startTime: number,
    endTime: number
}