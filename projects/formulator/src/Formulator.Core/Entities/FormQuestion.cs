using System.ComponentModel.DataAnnotations;

namespace Formulator.Core.Entities
{
    public class FormQuestion : BaseEntity
    {
        [Key]
        public int QuestionId { get; set; }

        public int FormId { get; set; }
        public Form Form { get; set; } = null!;

        public required string Text { get; set; }
        public string? Description { get; set; }

        public bool Required { get; set; } = false;
    }
}