using System.ComponentModel.DataAnnotations;

namespace Formulator.Core.Entities
{
    public class FormSubmission : BaseEntity
    {
        [Key]
        public int SubmissionId { get; set; }

        public int FormId { get; set; }
        public Form Form { get; set; } = null!;

        public int? SubmittorId { get; set; }
        public User Submittor { get; set; } = null!;

        // todo: review proper way of doing this
        public ICollection<FormSubmissionResponse> Responses { get; } = new List<FormSubmissionResponse>();
    }
}
