// import { LinksFunction } from "@remix-run/node";

import type React from "react";
import { useRef, useState } from "react";
import { Cropper, type ReactCropperElement } from "react-cropper";
// import "cropperjs/dist/cropper.css";

// export const clientLoader = () => {
//   return null;
// };

/**
 * 画像切り取りのサンプル
 */
const MyCropper: React.FC<{
  /**
   * file オブジェクトから解析した src ソース
   */
  readonly src: string | undefined;
  /**
   * 切り取り結果を受け取るコールバック
   */
  readonly resultBase64: (s: string) => void;
}> = ({ src, resultBase64 }) => {
  const cropperRef = useRef<ReactCropperElement>(null);

  const onCrop = () => {
    const cropper = cropperRef.current?.cropper;
    const r = cropper?.getCroppedCanvas().toDataURL();

    if (r === undefined) {
      return;
    }
    // コールバックに返却
    resultBase64(r);
  };

  if (src === undefined) {
    return <></>;
  }

  return (
    <div>
      <Cropper
        src={src}
        style={{ height: 400, width: "100%" }}
        // Cropper.js options
        initialAspectRatio={16 / 9}
        // aspectRatio={1 / 1}
        guides={true}
        crop={onCrop}
        ref={cropperRef}
      />
    </div>
  );
};

/**
 * input タグで読み取った画像のパスを取り出します
 *
 * @returns src url
 */
const LoadImageFromInput = (e: React.ChangeEvent<HTMLInputElement>): string | undefined => {
  const file = e.target.files?.[0];
  // console.log(file);

  if (file === undefined) {
    return undefined;
  }

  const blobUrl = window.URL.createObjectURL(file);
  return blobUrl;
};

const Demo: React.FC = () => {
  // 読み込んだファイルパス
  const [src, setSrc] = useState<string>();
  // 切り取り結果
  const [resultBase64, SetResultBase64] = useState<string>();

  return (
    <div>
      <input
        accept="image/jpeg, image/png"
        onChange={(e) => {
          setSrc(LoadImageFromInput(e));
        }}
        type="file"
      />
      {/* 切り取りコンポーネント */}
      <MyCropper src={src} resultBase64={SetResultBase64} />
      {/* 切り取り結果の表示 */}
      <img width={400} src={resultBase64} alt="cropping result" />
    </div>
  );
};

const Page = () => {
  return <Demo />;
};

export default Page;
