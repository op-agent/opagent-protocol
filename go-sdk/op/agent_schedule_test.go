package op

import "testing"

func TestScheduleValidateRejectsMultipleModes(t *testing.T) {
	err := (Schedule{
		Cron:  "0 9 * * *",
		Every: "1h",
	}).Validate()
	if err == nil {
		t.Fatal("Validate() error = nil, want error")
	}
}

func TestScheduleValidateAcceptsTime(t *testing.T) {
	if err := (Schedule{Time: "09:00"}).Validate(); err != nil {
		t.Fatalf("Validate() error = %v, want nil", err)
	}
}

func TestRunValidateRejectsScheduleWithoutScheduledLifecycle(t *testing.T) {
	err := (Run{
		Lifecycle: LifecycleOnDemand,
		Schedule:  Schedule{Every: "1h"},
	}).Validate()
	if err == nil {
		t.Fatal("Validate() error = nil, want error")
	}
}

func TestRunValidateRequiresScheduleForScheduledLifecycle(t *testing.T) {
	err := (Run{
		Lifecycle: LifecycleScheduled,
	}).Validate()
	if err == nil {
		t.Fatal("Validate() error = nil, want error")
	}
}
