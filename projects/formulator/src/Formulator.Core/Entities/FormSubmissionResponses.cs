namespace Formulator.Core.Entities
{
    public class FormSubmissionResponse : BaseEntity
    {
        public int SubmissionId { get; set; }
        public FormSubmission Submission { get; set; } = null!;

        public int QuestionId { get; set; }
        public FormQuestion Question { get; set; } = null!;

        public required string Value { get; set; }
    }
}
