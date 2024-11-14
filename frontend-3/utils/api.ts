export function uploadWithPresignedURL({
  file,
  presignedURL,
  onProgress,
  onError,
}: {
  file: File;
  presignedURL: string;
  onProgress: (progress: number) => void;
  onError: (error: number) => void;
}) {
  return new Promise((resolve, reject) => {
    const xhr = new XMLHttpRequest();
    xhr.open("PUT", presignedURL);
    xhr.onload = () => {
      console.log("Upload completed successfully");
      resolve(xhr.status);
    };
    xhr.onerror = () => {
      console.error("Upload failed");
      onError(xhr.status);
      reject(xhr.status);
    };
    xhr.upload.onprogress = (event) => {
      if (event.lengthComputable) {
        const percentComplete = (event.loaded / event.total) * 100;
        onProgress(percentComplete);
      }
    };

    xhr.send(file);
  });
}
