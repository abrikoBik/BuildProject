export interface BlueprintProp {
    name: string;
    size: number;
    uuid: string;
}

export interface EditableCellProps {
    value: string;
    onSave?: (newValue: string) => void;
    onChange?: (newValue: string) => void;
    inputStyle?: string;
    textStyle?: string;
    textArea?: boolean;
}