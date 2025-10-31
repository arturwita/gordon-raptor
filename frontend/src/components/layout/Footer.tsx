import { memo, type FC } from "react";

const Footer: FC = () => (
  <footer className="absolute bottom-4 text-sm text-gray-500 dark:text-gray-400">
    © {new Date().getFullYear()} Gordon Raptor. All rights reserved.
  </footer>
);

export default memo(Footer);
