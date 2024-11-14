import { MathFloor } from "@/utilities/utilities";

/**
 * 二つの日時データの差を、日・時・分・秒に分解して求める。
 */
export const PlatDateFormatter = (args: { platTimestamp: Date; currentTime: Date }): string => {
  const diff = args.currentTime.getTime() - args.platTimestamp.getTime();

  if (diff <= 0.0) {
    return "現在";
  }

  const seconds = MathFloor(diff / 1000);
  if (seconds < 60) {
    return `${seconds.toString()} 秒前`;
  }

  const minutes = MathFloor(seconds / 60);
  if (minutes < 60) {
    return `${minutes.toString()} 分前`;
  }

  const hours = MathFloor(minutes / 60);
  if (hours < 24) {
    return `${hours.toString()} 時間前`;
  }

  const days = MathFloor(hours / 24);
  if (days < 30) {
    return `${days.toString()} 日前`;
  }

  const months = MathFloor(days / 30);
  if (months < 12) {
    return `${months.toString()} ヶ月前`;
  }

  const years = MathFloor(months / 12);
  return `${years.toString()} 年前`;
};
