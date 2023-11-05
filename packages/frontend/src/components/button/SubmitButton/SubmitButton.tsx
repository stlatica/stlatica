export type SubmitButtonProps = {
  children: React.ReactNode;
  onClick?: () => void;
  className?: string | undefined;
};

/**
 *
 */
export const SubmitButton: React.FC<SubmitButtonProps> = ({ children, onClick, className }) => {
  return (
    <button
      className={
        "text-blue-900 bg-white border border-blue-300 focus:outline-none hover:bg-blue-100 focus:ring-4 focus:ring-blue-200 font-medium rounded-lg text-sm px-5 py-2.5 mr-2 mb-2 dark:bg-blue-600 dark:text-white dark:border-blue-600 dark:hover:bg-blue-700 dark:hover:border-blue-600 " +
        className
      }
      onClick={onClick}
      type="button"
    >
      {children}
    </button>
  );
};
