import { type FC, memo } from "react";

interface TestProps {}

const Test: FC<TestProps> = () => {
  return <div>test</div>;
};

export default memo(Test);
