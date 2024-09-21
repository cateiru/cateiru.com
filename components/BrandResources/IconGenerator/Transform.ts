export class Transform {
  private context: CanvasRenderingContext2D | null;

  private image: HTMLImageElement | undefined;

  private afterWidth: number | undefined;
  private afterHeight: number | undefined;

  constructor(width?: number, height?: number) {
    this.context = document.createElement("canvas").getContext("2d");

    this.afterWidth = width;
    this.afterHeight = height;
  }

  /**
   * 画像を読み込む
   * 同一オリジン以外からの読み込みはCORSでエラーが発生するので注意
   *
   * @param {string} src - 画像をロードする。CORSがあるので同一オリジンのみ可
   */
  public async loadImageFromSrc(src: string) {
    this.image = await loadImage(src);
    this.draw();
  }

  /**
   * 透過背景を別の色に置き換える
   *
   * @param {number} r - red
   * @param {number} g - green
   * @param {number} b - blue
   */
  public setBg(r: number, g: number, b: number) {
    if (!this.context) throw new Error("context is null");
    if (typeof this.image === "undefined")
      throw new Error("image is undefined");

    const imageData = this.context.getImageData(
      0,
      0,
      this.context.canvas.width,
      this.context.canvas.height,
    );

    for (let i = 0; i < imageData.width * imageData.height; i++) {
      if (imageData.data[i * 4 + 3] === 0) {
        imageData.data[i * 4] = r;
        imageData.data[i * 4 + 1] = g;
        imageData.data[i * 4 + 2] = b;
        imageData.data[i * 4 + 3] = 255;
      }
    }
    this.context.putImageData(imageData, 0, 0);
  }

  public setCircle() {
    if (!this.context) throw new Error("context is null");
    if (typeof this.image === "undefined")
      throw new Error("image is undefined");

    this.context.globalCompositeOperation = "destination-in";

    const size = this.context.canvas.width / 2;

    this.context.arc(
      this.context.canvas.width / 2,
      this.context.canvas.height / 2,
      size,
      0,
      Math.PI * 2,
    );
    this.context.fill();
  }

  private draw() {
    if (!this.context) throw new Error("context is null");
    if (typeof this.image === "undefined")
      throw new Error("image is undefined");

    const w = this.afterWidth ?? this.image.naturalWidth;
    const h = this.afterHeight ?? this.image.naturalHeight;

    this.context.canvas.width = w;
    this.context.canvas.height = h;

    this.context.drawImage(
      this.image,
      0,
      0,
      this.image.naturalWidth,
      this.image.naturalHeight,
      0,
      0,
      w,
      h,
    );
  }

  /**
   * 画像を出力する
   *
   * @param {string} type - 画像タイプ。 `image/png`, `image/jpeg`など
   * @param {number} quality - クオリティ。jpegかwebpの場合のみ。0~1
   * @returns {Blob} - 画像データ
   */
  public async export(type: string, quality = 1.0): Promise<Blob | null> {
    return await new Promise<Blob | null>((resolve) => {
      if (!this.context) return null;
      this.context.canvas.toBlob(resolve, type, quality);
    });
  }
}

/**
 * 画像の幅と高さを取得する
 *
 * @param {File} b - image file object
 * @returns {[number, number]} - [幅, 高さ]
 */
export async function getSize(b: File): Promise<[number, number]> {
  const image = await loadImage(URL.createObjectURL(b));

  return [image.width, image.height];
}

export const loadImage = async (src: string): Promise<HTMLImageElement> =>
  await new Promise((resolve, reject) => {
    const image = new Image();
    image.addEventListener("load", () => resolve(image));
    image.addEventListener("error", reject);
    image.src = src;
  });
