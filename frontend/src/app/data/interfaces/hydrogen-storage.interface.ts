export interface HydrogenStorage {
    id: number
    temperature_range: {
        Lower: number;
        Upper: number;
    }
    nominal_voltage: number;
    recorded_at: string;
    brand_name: string;
    model_name: string;
    specific_energy: number;
    efficiency: number;
}