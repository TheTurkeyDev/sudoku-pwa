import styles from './loading-spinner.module.css';

export const LoadingSpinner = () => (
    <div class={styles.LoadingSpinner} />
);

export const CenteredLoadingSpinner = () => (
    <div class={styles.CenteredLoadingSpinner}>
        <LoadingSpinner />
    </div>
);