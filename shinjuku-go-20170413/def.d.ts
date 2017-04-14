declare module Submarine.Battle {
  interface Point {
    x: number;
    y: number;
  }

  // リアルタイムメッセージは TyphenApi.RealTimeMessage を継承
  interface Ping extends TyphenApi.RealTimeMessage {
    message: string;
  }

  // Web サーバーのエンドポイントの定義は function を使用
  function ping(message: string): { message: string; };
}
