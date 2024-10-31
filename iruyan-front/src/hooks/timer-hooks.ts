import { useState, useEffect, useRef } from "react";

type TimerMode = "work" | "break";

interface UseTimerProps {
  initialWorkTime?: number; // 秒単位
  initialBreakTime?: number; // 秒単位
}

export function useTimer({
  initialWorkTime = 1500,
  initialBreakTime = 300,
  onTick,
}: UseTimerProps & { onTick?: () => void }) {
  const [timeLeft, setTimeLeft] = useState(initialWorkTime);
  const [isRunning, setIsRunning] = useState(false);
  const [mode, setMode] = useState<TimerMode>("work");
  const [totalWorkTime, setTotalWorkTime] = useState(initialWorkTime);
  const [totalBreakTime, setTotalBreakTime] = useState(initialBreakTime);
  const timerId = useRef<NodeJS.Timeout | null>(null);

  // タイマーの動作を管理
  useEffect(() => {
    if (isRunning) {
      timerId.current = setInterval(() => {
        setTimeLeft((prevTime) => {
          if (prevTime <= 1) {
            clearInterval(timerId.current!);
            setIsRunning(false);
            // タイマー終了時に自動でモードを切り替える場合は以下を有効化
            // switchMode();
            return 0;
          }
          return prevTime - 1;
        });
        if (onTick) onTick();
      }, 1000);
    } else if (timerId.current !== null) {
      clearInterval(timerId.current);
    }
    return () => {
      if (timerId.current !== null) {
        clearInterval(timerId.current);
      }
    };
  }, [isRunning, onTick]);

  const reset = () => {
    setTimeLeft(mode === "work" ? totalWorkTime : totalBreakTime);
    setIsRunning(false);
  };

  const stop = () => {
    setIsRunning(false);
  };

  const start = () => {
    setIsRunning(true);
  };

  const switchMode = () => {
    const newMode = mode === "work" ? "break" : "work";
    setMode(newMode);
    const newTime = newMode === "work" ? totalWorkTime : totalBreakTime;
    setTimeLeft(newTime);
    setIsRunning(false);
  };

  const adjustTime = (adjustment: number) => {
    if (mode === "work") {
      setTotalWorkTime((prevTotal) => {
        const newTotal = Math.max(300, prevTotal + adjustment);
        if (!isRunning) {
          setTimeLeft(newTotal);
        }
        return newTotal;
      });
    } else {
      setTotalBreakTime((prevTotal) => {
        const newTotal = Math.max(300, prevTotal + adjustment);
        if (!isRunning) {
          setTimeLeft(newTotal);
        }
        return newTotal;
      });
    }
  };

  const totalTime = mode === "work" ? totalWorkTime : totalBreakTime;
  const progress = ((totalTime - timeLeft) / totalTime) * 100;

  return {
    timeLeft,
    isRunning,
    mode,
    totalTime,
    progress,
    reset,
    stop,
    start,
    switchMode,
    adjustTime,
  };
}
