"use client";

import { useEffect } from "react";

const STORAGE_KEY = "cateiru-scroll-hint-shown";

export const ScrollHint = () => {
  useEffect(() => {
    if (localStorage.getItem(STORAGE_KEY) === "1") {
      return;
    }
    if (window.matchMedia("(prefers-reduced-motion: reduce)").matches) {
      return;
    }

    const timers: number[] = [];
    let cancelled = false;

    const cancel = (event: Event) => {
      if (!event.isTrusted) return;
      if (cancelled) return;
      cancelled = true;
      for (const id of timers) {
        clearTimeout(id);
      }
      localStorage.setItem(STORAGE_KEY, "1");
      removeCancelListeners();
    };

    const addCancelListeners = () => {
      window.addEventListener("wheel", cancel, { passive: true });
      window.addEventListener("touchstart", cancel, { passive: true });
    };

    const removeCancelListeners = () => {
      window.removeEventListener("wheel", cancel);
      window.removeEventListener("touchstart", cancel);
    };

    timers.push(
      window.setTimeout(() => {
        if (cancelled) return;
        addCancelListeners();
        window.scrollTo({ top: window.innerHeight * 0.1, behavior: "smooth" });

        timers.push(
          window.setTimeout(() => {
            if (cancelled) return;
            window.scrollTo({ top: 0, behavior: "smooth" });

            timers.push(
              window.setTimeout(() => {
                if (cancelled) return;
                localStorage.setItem(STORAGE_KEY, "1");
                removeCancelListeners();
              }, 900)
            );
          }, 1200)
        );
      }, 500)
    );

    return () => {
      for (const id of timers) {
        clearTimeout(id);
      }
      removeCancelListeners();
    };
  }, []);

  return null;
};
