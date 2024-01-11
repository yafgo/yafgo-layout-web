import dayjs from 'dayjs';

/**
 * 将 unix 时间戳格式化为 date
 * @param { number } time unix 时间戳
 * @returns
 */
export function TimestampToDate(time: number) {
    return time <= 0 ? '-' : dayjs(time * 1000).format('yyyy-MM-dd');
}

/**
 * 将 unix 时间戳格式化为 date
 * @param { number } time unix 时间戳
 * @param { boolean } isMS 传入的是否是毫秒时间戳, 默认是 unix 时间戳
 * @returns
 */
export function TimestampToDatetime(time: number | string, isMS = false) {
    if (typeof time === 'string') {
        return dayjs(time).format('YYYY-MM-DD HH:mm:ss');
    }
    return time <= 0 ? '' : dayjs(isMS ? time : time * 1000).format('yyyy-MM-dd HH:mm:ss');
}

/**
 * 将 date 格式化为 unix 时间戳
 * @param { string|number } time 时间字符串
 * @returns
 */
export function DatetimeToTimestamp(time: string | number) {
    const d = new Date(time)
    const ts = Math.round(d.getTime() / 1000)
    return time ? ts : 0;
}