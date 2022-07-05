import styles from './css/options.module.css';

export const Options = () => {
    return (
        <div class={styles.Options}>
            {
                Array.from(Array(9)).map((_, i) => <span>{i + 1}</span>)
            }
        </div>
    )
}