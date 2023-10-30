export type CancelButtonProps = {
  children: React.ReactNode;
  onClick?: () => void;
  className?: string | undefined;
};

/**
 *
 */
export const CancelButton: React.FC<CancelButtonProps> = ({ children, onClick, className }) => {
  return (
    <button
      className={
        "text-gray-900 bg-white border border-gray-300 focus:outline-none hover:bg-gray-100 focus:ring-4 focus:ring-gray-200 font-medium rounded-lg text-sm px-5 py-2.5 mr-2 mb-2 dark:bg-gray-800 dark:text-white dark:border-gray-600 dark:hover:bg-gray-700 dark:hover:border-gray-600 " +
        className
      }
      onClick={onClick}
      type="button"
    >
      {children}
    </button>
  );
};
