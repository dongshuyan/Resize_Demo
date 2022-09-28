#include <opencv2/opencv.hpp>
#include <opencv2/core/core.hpp>
#include <opencv2/highgui/highgui.hpp>
#include <opencv2/imgproc/imgproc.hpp>
#include <iostream>
using namespace cv;
using namespace std;
int main() {
    String img_path = "your_image_path"; //When using relative path, be relative to the EXECUTABLE file
    Mat img = imread(img_path);
    Size r_size = Size(64, 64);
    Mat resized_img;
    if (!img.data) {
        cout << "Image reading error !" << endl;
        return 1;
    }
    resize(img, resized_img, r_size);
    imwrite(img_path, resized_img);
    return 0;
}
