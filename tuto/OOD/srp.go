package main

type Report struct {
}

type FinanceReport struct {
}

func (f *FinanceReport) MakeReport() *Report {
	// ...
	// FinanceReport의 책임 1
	return &Report{}
}

/*
func (f *FinanceReport) SendReport(email string) {
	// FinanceReport의 책임 2 >> 단일 책임 원칙에 어긋남, 그러므로 send 기능을 구분 지을 필요 있음

	switch {
	case 1: email ...
	case 2: file ...
	case 3: http ...
	case 4: network ...
	} // 이처럼 책임을 여러개 갖게 되면 변경 사항이 생길 때마다 매번 수정해야함
}
*/

// ReportSender interface를 만듦으로써 개별 작업의 확장을 간편하게 함, 또한 변경에 닫혀 있음 (S, O 모두 만족)
type ReportSender interface {
	SendReport(*Report)
}

// email 을 보내는 객체
type EmailReportSender struct {
}

func (s *EmailReportSender) SendReport(r *Report) {
	// email ...
}

// File 을 보내는 객체
type FileReportSender struct {
}

func (s *FileReportSender) SendReport(r *Report) {
	// file ...
}

func main() {

}
