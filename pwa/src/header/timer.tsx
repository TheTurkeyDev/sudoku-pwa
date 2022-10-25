import { createSignal, onCleanup } from 'solid-js';
import { useSudoku } from '../context/sudoku';

const timeBetweenDates = (from: number, to: number) => {
    const difference = to - from;

    const timeData = {
        minutes: '00',
        seconds: '00'
    };

    if (difference > 0) {
        const secondsRaw = Math.floor(difference / 1000);
        const minutes = (Math.floor(secondsRaw / 60)) % 60;
        const seconds = secondsRaw % 60;

        timeData.minutes = minutes < 10 ? `0${minutes}` : `${minutes}`;
        timeData.seconds = seconds < 10 ? `0${seconds}` : `${seconds}`;
    }
    return {
        timeData,
        difference
    };
};


export const Timer = () => {
    const { startTime, endTime, gameState } = useSudoku();
    const [timerDetails, setTimerDetails] = createSignal(timeBetweenDates(startTime(), new Date().getTime()).timeData);

    const timer = setInterval(() => {
        setTimerDetails(timeBetweenDates(startTime(), gameState() === 1 ? new Date().getTime() : endTime()).timeData);
    }, 1000);

    onCleanup(() => clearInterval(timer));

    return (
        <span>{timerDetails().minutes}:{timerDetails().seconds}</span>
    );
};