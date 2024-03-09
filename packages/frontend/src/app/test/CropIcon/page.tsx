"use client";
import { Button } from "@mantine/core";
import React, { useState, useCallback } from "react";
import Cropper, { Area } from "react-easy-crop";

import getCroppedImg from "./getCroppedImg";

// 参考
// ファイルのアップロード：https://zenn.dev/yuyan/articles/f35da08770a135
// Cropper：https://codesandbox.io/p/sandbox/suspicious-http-rqfrgb?file=%2Fsrc%2FApp.tsx%3A163%2C9&from-embed=
export default function Home() {
  const [base64Image, setBase64Image] = useState<string>(""); // 画像を一枚だけ保持するためのstate

  const handleInputFile = (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files ? e.target.files[0] : null; // 単一のファイルを選択
    if (!file) {
      return;
    }

    const reader = new FileReader();
    reader.onloadend = () => {
      const result = reader.result;
      if (typeof result === "string") {
        setBase64Image(result); // 前回の画像を上書きして新しい画像をセット
      }
    };
    reader.readAsDataURL(file); // 画像ファイルをbase64形式で読み込む
  };

  /** 画像の拡大縮小倍率 */
  const [zoom, setZoom] = useState<number>(1);

  /** 切り取る領域の情報 */
  const [crop, setCrop] = useState({ x: 0, y: 0 });
  /** 切り取る領域の情報 */
  const [croppedAreaPixels, setCroppedAreaPixels] = useState<Area>();

  /** 切り取ったあとの画像URL */
  const [croppedImgSrc, setCroppedImgSrc] = useState("");

  /**
   * 切り取り完了後、切り取り領域の情報をセット
   */
  const onCropComplete = useCallback((croppedArea: Area, croppedAreaPixels: Area) => {
    setCroppedAreaPixels(croppedAreaPixels);
  }, []);

  /**
   * 切り取り後の画像を生成し画面に表示
   */
  const showCroppedImage = useCallback(async () => {
    if (!croppedAreaPixels) return;
    try {
      const croppedImage = await getCroppedImg(base64Image, croppedAreaPixels);
      setCroppedImgSrc(croppedImage);
    } catch (e) {
      console.error(e);
    }
  }, [croppedAreaPixels, base64Image]);

  return (
    <div>
      <div className="relative h-[300px] w-[300px]">
        <Cropper
          crop={crop}
          cropShape="round"
          cropSize={{ width: 300, height: 300 }}
          image={base64Image}
          maxZoom={5}
          minZoom={1}
          onCropChange={setCrop}
          onCropComplete={onCropComplete}
          onZoomChange={setZoom}
          showGrid={false}
          zoom={zoom}
        />
      </div>
      <div>
        <input
          aria-labelledby="Zoom"
          max={5}
          min={1}
          onChange={(e) => {
            setZoom(e.target.valueAsNumber);
          }}
          step={0.1}
          type="range"
          value={zoom}
        />
      </div>
      <div>
        <input accept="image/jpeg, image/png" onChange={handleInputFile} type="file" />
        <Button
          onClick={() => {
            showCroppedImage();
          }}
        >
          OK
        </Button>
      </div>
      <div>
        {croppedImgSrc ? (
          <img alt="Cropped" className="h-[300px] w-[300px]" src={croppedImgSrc} />
        ) : (
          <div>The cropped image is displayed here</div>
        )}
      </div>
    </div>
  );
}
