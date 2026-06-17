// Tipe Data Slice
// * Tipe Data Slice adalah potongan dari data Array
// * Slice mirip dengan Array, yang membedakan adalah ukuran Slice bisa berubah
// * Slice dan Array terkoneksi, dimana Slice adalah data yang mengakses sebagian
//   atau seluruh data di array

// Detail Tipe Data Slice
// * Tipe data slice memiliki 3 data, yaitu pointer, length, dan capacity
// * Pointer adalah penunjuk data pertama di array para slice
// * Length adalah panjang dari slice
// * Capacity adalah kapasitas dari slice, dimana length tidak boleh lebih dari capacity

package main

import "fmt"

func main(){
	var bulan = [12]string {
		"januari",
		"februari",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}
	fmt.Println(bulan[0:2])
	fmt.Println(bulan[1])

	//Builtin function
	fmt.Println("======Builtin Function=======")
	var slices = bulan[0:8]
	fmt.Println(len(bulan), "Len Bulan")
	fmt.Println(len(slices), "Len Slices")
	fmt.Println(cap(slices), "Cap Slices")

	// jika value bulan diubah maka slices juga terubah (Mutable)
	// begitupun jika slices diubah valuenya maka value dibulan juga terubah
	// fmt.Println(slices[0:2], "Sebelum diubah")
	// bulan[0] = "janjan"
	// fmt.Println(slices[0:2], "Setelah diubah")

	fmt.Println("====== Append Function========")
	var slices2 = bulan[10:]
	fmt.Println(slices2)
	
	var slices3 = append(slices2, "imron")
	fmt.Println(slices3)
	slices3[1] = "bukan desember"
	fmt.Println(slices3)
	fmt.Println(slices2)
	fmt.Println(bulan)

	fmt.Println("====== Make Function========")
	newSlices := make([]string, 2, 5) // (tipe_data_array, maksimal_data, capacity)
	newSlices[0] = "imron"
	newSlices[1] = "syabani"
	fmt.Println(newSlices)
	fmt.Println(len(newSlices))
	fmt.Println(cap(newSlices))
	newSlices2 := append(newSlices, "imronsyabani", "hallo imron")
	fmt.Println(newSlices2, "hasil append newslices")
	newSlices2[3] = "imronsyabani"
	fmt.Println(newSlices2, "setelah di reassign valuenya")
	fmt.Println(newSlices, "hasil dari newslices")
	newSlices2[1] = "hallo syabani" //jadinya mutable newSlices terubah newSlices2 juga terubah
	fmt.Println(newSlices2, "setelah di reassign dengan index yang ada pada newslices valuenya")
	fmt.Println(newSlices, "hasil dari newslices")
	
	fmt.Println("====== Copy Function========")
	copySlices := make([]string, len(newSlices), cap(newSlices))
	copy(copySlices, newSlices)
	copySlices[1] = "syabani"
	fmt.Println(copySlices, "Ini copy Array")
	fmt.Println(newSlices, "ini slices")
}