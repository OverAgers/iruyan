export const getCurrentTimestamp = (): number => {
    return Math.floor(Date.now() / 1000);
};
export const formatDate = (timestamp: number): string => {
    return new Date(timestamp * 1000).toLocaleDateString('ja-JP');
};