import styles from './header.module.css';
import { useSudoku } from '../context/sudoku';
import { Timer } from './timer';

const mapDifficultyToText = (dif: number) => {
    switch (dif) {
        case 0:
            return 'Beginner';
        case 1:
            return 'Easy';
        case 2:
            return 'Medium';
        case 3:
            return 'Hard';
        default:
            return 'Puzzle';
    }
};
export const Header = () => {
    const { difficulty, startTime } = useSudoku();

    return (
        <div class={styles.HeaderContent}>
            <span>{mapDifficultyToText(difficulty())}</span>
            <span>-</span>
            <Timer />
        </div>
    );
};